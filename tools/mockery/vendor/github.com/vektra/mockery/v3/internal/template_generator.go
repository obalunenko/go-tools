package internal

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"go/format"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/vektra/mockery/v3/config"
	"github.com/vektra/mockery/v3/internal/file"
	"github.com/vektra/mockery/v3/internal/stackerr"
	"github.com/vektra/mockery/v3/template"
	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

type Formatter string

const (
	FormatGofmt     Formatter = "gofmt"
	FormatGoImports Formatter = "goimports"
	FormatNoop      Formatter = "noop"
)

var (
	//go:embed mock_matryer.templ
	templateMatryer string
	//go:embed mock_matryer.templ.schema.json
	templateMatryerJSONSchema string
	//go:embed mock_testify.templ
	templateTestify string
	//go:embed mock_testify.templ.schema.json
	templateTestifyJSONSchema string
)

var errBadHTTPStatus = errors.New("failed to download file")

var styleTemplates = map[string]string{
	"matryer": templateMatryer,
	"testify": templateTestify,
}

var jsonSchemas = map[string]string{
	"matryer": templateMatryerJSONSchema,
	"testify": templateTestifyJSONSchema,
}

// findPkgPath returns the fully-qualified go import path of a given dir. The
// dir must be relative to a go.mod file. In the case it isn't, an error is returned.
func findPkgPath(dirPath string) (string, error) {
	dirPath, err := file.CleanPath(dirPath)
	if err != nil {
		return "", stackerr.NewStackErr(err)
	}
	goModFile, content, err := file.FindInHierarchy(dirPath, []string{"go.mod"})
	if err != nil {
		return "", stackerr.NewStackErr(err)
	}
	dirRelative, err := filepath.Rel(filepath.FromSlash(filepath.Dir(goModFile)), filepath.FromSlash(dirPath))
	if err != nil {
		return "", stackerr.NewStackErr(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(content))
	// Iterate over each line
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "module") {
			continue
		}
		moduleName := strings.Split(scanner.Text(), "module ")[1]
		return filepath.ToSlash(filepath.Clean(filepath.Join(moduleName, dirRelative))), nil
	}
	return "", stackerr.NewStackErr(ErrGoModInvalid)
}

type TemplateGenerator struct {
	formatter           Formatter
	inPackage           bool
	requireSchemaExists bool
	registry            *template.Registry
	templateName        string
	templateSchema      string
	pkgConfig           *config.Config
	pkgName             string
	remoteTemplateCache map[string]*RemoteTemplate
}

func NewTemplateGenerator(
	ctx context.Context,
	srcPkg *packages.Package,
	outPkgFSPath string,
	templateName string,
	templateSchema string,
	requireSchemaExists bool,
	remoteTemplateCache map[string]*RemoteTemplate,
	formatter Formatter,
	pkgConfig *config.Config,
	pkgName string,
	forceInPackage *bool,
) (*TemplateGenerator, error) {
	log := *zerolog.Ctx(ctx)
	var err error

	srcPkgFSPath := filepath.Dir(srcPkg.GoFiles[0])
	if !filepath.IsAbs(outPkgFSPath) {
		outPkgFSPath, err = filepath.Abs(outPkgFSPath)
		if err != nil {
			log.Err(err).Msg("failed to make absolute path")
			return nil, stackerr.NewStackErr(err)
		}
	}
	if !filepath.IsAbs(srcPkgFSPath) {
		srcPkgFSPath, err = filepath.Abs(srcPkgFSPath)
		if err != nil {
			log.Err(err).Msg("failed to make absolute path")
			return nil, stackerr.NewStackErr(err)
		}
	}
	srcPkgFSPath = filepath.ToSlash(filepath.Clean(srcPkgFSPath))
	outPkgFSPath = filepath.ToSlash(filepath.Clean(outPkgFSPath))

	newLogger := zerolog.Ctx(ctx).With().
		Str("srcPkgFSPath", srcPkgFSPath).
		Str("outPkgFSPath", outPkgFSPath).
		Str("src-pkg-name", srcPkg.Name).
		Str("out-pkg-name", pkgName).
		Logger()
	log = newLogger

	outPkgPath, err := findPkgPath(outPkgFSPath)
	if err != nil {
		log.Err(err).Msg("failed to find output package path")
		return nil, err
	}
	log = log.With().Str("outPkgPath", outPkgPath).Logger()

	var inPackage bool
	if forceInPackage != nil {
		log.Debug().Bool("inpackage", *forceInPackage).Msg("inpackage value provided, forcing to value")
		inPackage = *forceInPackage
	} else {
		// Note: Technically, go allows test files to have a different package name
		// than non-test files. In this case, the test files have to import the source
		// package just as if it were in a different directory.
		if pkgName == srcPkg.Name && srcPkgFSPath == outPkgFSPath {
			log.Debug().Msg("output package detected to be in-package of original package")
			inPackage = true
		} else {
			log.Debug().Msg("output package detected to not be in-package of original package")
		}
	}

	reg, err := template.NewRegistry(srcPkg, outPkgPath, inPackage)
	if err != nil {
		return nil, fmt.Errorf("creating new registry: %w", err)
	}

	return &TemplateGenerator{
		templateName:        templateName,
		templateSchema:      templateSchema,
		requireSchemaExists: requireSchemaExists,
		registry:            reg,
		formatter:           formatter,
		inPackage:           inPackage,
		pkgConfig:           pkgConfig,
		pkgName:             pkgName,
		remoteTemplateCache: remoteTemplateCache,
	}, nil
}

func (g *TemplateGenerator) format(src []byte) ([]byte, error) {
	switch g.formatter {
	case FormatGoImports:
		return goimports(src)
	case FormatGofmt:
		return gofmt(src)
	case FormatNoop:
		return src, nil
	}

	return nil, fmt.Errorf("unknown formatter type: %s", g.formatter)
}

func getTypePath(t types.Type) (paramPkgPath string, paramObjName string) {
	switch t := t.(type) {
	case *types.Named:
		pkg := t.Obj().Pkg()
		if pkg != nil {
			paramPkgPath = pkg.Path()
		}
		paramObjName = t.Obj().Name()
	case *types.Alias:
		pkg := t.Obj().Pkg()
		if pkg != nil {
			paramPkgPath = pkg.Path()
		}
		paramObjName = t.Obj().Name()
	case *types.Pointer:
		paramPkgPath, paramObjName = getTypePath(t.Elem())
	}
	return paramPkgPath, paramObjName
}

func (g *TemplateGenerator) methodData(ctx context.Context, method *types.Func, ifaceConfig *config.Config) (template.Method, error) {
	log := zerolog.Ctx(ctx)

	methodScope := g.registry.MethodScope()

	signature := method.Type().(*types.Signature)
	params := make([]template.Param, signature.Params().Len())

	for j := 0; j < signature.Params().Len(); j++ {
		param := signature.Params().At(j)
		log.Debug().Str("param-string", param.String()).Msg("found parameter")
		for _, imprt := range g.registry.Imports() {
			log.Debug().Str("import", imprt.Path()).Str("import-qualifier", imprt.Qualifier()).Msg("existing imports")
		}

		paramPkgPath, paramObjName := getTypePath(param.Type())

		replacement := ifaceConfig.GetReplacement(paramPkgPath, paramObjName)
		if replacement != nil {
			log.Debug().Str("replace-to-pkg-path", replacement.PkgPath).Str("replace-to-type-name", replacement.TypeName).Msg("found replacement")
		} else {
			log.Debug().Str("param-pkg-path", paramPkgPath).Msg("replacement not found")
		}
		v, err := methodScope.AddVar(ctx, param, "", replacement)
		if err != nil {
			return template.Method{}, err
		}
		params[j] = template.Param{
			Var:      v,
			Variadic: signature.Variadic() && j == signature.Params().Len()-1,
		}
	}

	returns := make([]template.Param, signature.Results().Len())
	for j := 0; j < signature.Results().Len(); j++ {
		param := signature.Results().At(j)

		paramLog := log.With().Str("param-string", param.String()).Logger()
		paramCtx := paramLog.WithContext(ctx)
		paramLog.Debug().Msg("found return")

		paramPkgPath, paramObjName := getTypePath(param.Type())

		replacement := ifaceConfig.GetReplacement(paramPkgPath, paramObjName)
		if replacement != nil {
			log.Debug().Str("replace-to-pkg-path", replacement.PkgPath).Str("replace-to-type-name", replacement.TypeName).Msg("found replacement")
		} else {
			log.Debug().Str("param-pkg-path", paramPkgPath).Msg("replacement not found")
		}
		v, err := methodScope.AddVar(paramCtx, param, "", replacement)
		if err != nil {
			return template.Method{}, err
		}
		returns[j] = template.Param{
			Var:      v,
			Variadic: false,
		}
	}
	return template.Method{
		Name:    method.Name(),
		Params:  params,
		Returns: returns,
		Scope:   methodScope,
	}, nil
}

func explicitConstraintType(typeParam *types.Var) (t types.Type) {
	underlying := typeParam.Type().Underlying().(*types.Interface)
	// check if any of the embedded types is either a basic type or a union,
	// because the generic type has to be an alias for one of those types then
	for j := 0; j < underlying.NumEmbeddeds(); j++ {
		t := underlying.EmbeddedType(j)
		switch t := t.(type) {
		case *types.Basic:
			return t
		case *types.Union: // only unions of basic types are allowed, so just take the first one as a valid type constraint
			return t.Term(0).Type()
		}
	}
	return nil
}

func (g *TemplateGenerator) typeParams(ctx context.Context, tparams *types.TypeParamList) ([]template.TypeParam, error) {
	var tpd []template.TypeParam
	if tparams == nil {
		return tpd, nil
	}

	tpd = make([]template.TypeParam, tparams.Len())

	scope := g.registry.MethodScope()
	for i := 0; i < len(tpd); i++ {
		tp := tparams.At(i)
		typeParam := types.NewParam(token.Pos(i), tp.Obj().Pkg(), tp.Obj().Name(), tp.Constraint())
		v, err := scope.AddVar(ctx, typeParam, "", nil)
		if err != nil {
			return nil, err
		}
		tpd[i] = template.TypeParam{
			Param:      template.Param{Var: v},
			Constraint: explicitConstraintType(typeParam),
		}
	}

	return tpd, nil
}

// getTemplate returns the requested template and associated schema (if available).
func (g *TemplateGenerator) getTemplate(ctx context.Context) (string, *gojsonschema.Schema, error) {
	log := zerolog.Ctx(ctx).With().Str("template", g.templateName).Str("schema", g.templateSchema).Logger()
	ctx = log.WithContext(ctx)

	for _, protocol := range []string{"file://", "https://", "http://"} {
		var schema *gojsonschema.Schema
		var err error

		if !strings.HasPrefix(g.templateName, protocol) {
			continue
		}
		var remoteTemplate *RemoteTemplate
		if cachedRemoteTemplate, ok := g.remoteTemplateCache[g.templateName]; !ok {
			remoteTemplate = NewRemoteTemplate(g.templateName, g.templateSchema)
			g.remoteTemplateCache[g.templateName] = remoteTemplate
		} else {
			remoteTemplate = cachedRemoteTemplate
		}

		templateString, err := remoteTemplate.Template(ctx)
		if err != nil {
			log.Error().Msg("could not download template")
			return "", nil, fmt.Errorf("downloading template: %w", err)
		}
		if g.requireSchemaExists {
			schema, err = remoteTemplate.Schema(ctx)
			if err != nil {
				log.Error().Msg("could not get JSON schema")
				return "", nil, fmt.Errorf("downloading schema: %w", err)
			}
		}

		return templateString, schema, nil
	}

	// Embedded templates
	var styleExists bool
	templateString, styleExists := styleTemplates[g.templateName]
	if !styleExists {
		return "", nil, fmt.Errorf("template '%s' does not exist", g.templateName)
	}
	schema, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(jsonSchemas[g.templateName]))
	if err != nil {
		return "", nil, fmt.Errorf("generating schema: %w", err)
	}
	return templateString, schema, nil
}

func validateSchema(ctx context.Context, data template.Data, schema *gojsonschema.Schema) error {
	if schema == nil {
		return errors.New("jschema argument can't be nil")
	}
	if err := data.TemplateData.VerifyJSONSchema(ctx, schema); err != nil {
		return fmt.Errorf("validating template-data")
	}
	for _, intf := range data.Interfaces {
		if err := intf.TemplateData.VerifyJSONSchema(ctx, schema); err != nil {
			return fmt.Errorf("verifying template-data for %s: %w", intf.Name, err)
		}
	}
	return nil
}

func (g *TemplateGenerator) Generate(
	ctx context.Context,
	interfaces []*Interface,
) ([]byte, error) {
	log := zerolog.Ctx(ctx)
	log.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("template", g.templateName).Str("schema", g.templateSchema)
	})

	mockData := []template.Interface{}
	for _, ifaceMock := range interfaces {
		ifaceLog := log.With().
			Str("interface-name", ifaceMock.Name).
			Str("package-path", ifaceMock.Pkg.PkgPath).
			Str("mock-name", *ifaceMock.Config.StructName).
			Logger()
		ctx := ifaceLog.WithContext(ctx)

		ifaceLog.Debug().Msg("looking up interface in registry")
		iface, tparams, err := g.registry.LookupInterface(ifaceMock.Name)
		if err != nil {
			log.Err(err).Msg("error looking up interface")
			return []byte{}, err
		}
		ifaceLog.Debug().Msg("found interface")

		methods := make([]template.Method, iface.NumMethods())
		for i := 0; i < iface.NumMethods(); i++ {
			methodData, err := g.methodData(ctx, iface.Method(i), ifaceMock.Config)
			if err != nil {
				return nil, err
			}
			methods[i] = methodData
		}
		// Now that all methods have been generated, we need to resolve naming
		// conflicts that arise between variable names and package qualifiers.
		for _, method := range methods {
			method.Scope.ResolveVariableNameCollisions(
				zerolog.
					Ctx(ctx).
					With().
					Str("method-name", method.Name).
					Logger().
					WithContext(ctx))
		}

		ifaceLog.Debug().Str("template-data", fmt.Sprintf("%v", ifaceMock.Config.TemplateData)).Msg("printing template data")
		tParams, err := g.typeParams(ctx, tparams)
		if err != nil {
			return nil, err
		}
		mockData = append(mockData, template.NewInterface(
			ifaceMock.Name,
			*ifaceMock.Config.StructName,
			tParams,
			methods,
			ifaceMock.Config.TemplateData,
			template.NewComments(ifaceMock.TypeSpec, ifaceMock.GenDecl),
		))
	}

	data := template.NewData(
		g.pkgName, "", template.Packages{}, mockData, g.pkgConfig.TemplateData, g.registry,
	)
	if !g.inPackage {
		data.SrcPkgQualifier = g.registry.SrcPkgName() + "."
	}

	templateString, schema, err := g.getTemplate(ctx)
	if err != nil {
		log.Error().Msg("could not get template")
		return nil, fmt.Errorf("getting template: %w", err)
	}
	if schema != nil {
		if err := validateSchema(ctx, data, schema); err != nil {
			log.Error().Msg("failed to validate schema")
			return nil, fmt.Errorf("validating schema: %w", err)
		}
	}

	templ, err := template.New(templateString, g.templateName)
	if err != nil {
		return []byte{}, fmt.Errorf("creating new template: %w", err)
	}

	var buf bytes.Buffer
	log.Debug().Msg("executing template")
	if err := templ.Execute(&buf, data); err != nil {
		log.Error().Msg("failed to execute template")
		return []byte{}, fmt.Errorf("executing template: %w", err)
	}

	log.Debug().Msg("formatting file in-memory")
	formatted, err := g.format(buf.Bytes())
	if err != nil {
		scanner := bufio.NewScanner(strings.NewReader(buf.String()))
		for i := 1; scanner.Scan(); i++ {
			fmt.Printf("%d:\t%s\n", i, scanner.Text())
		}
		log.Err(err).Msg("can't format mock file in-memory")
		return []byte{}, fmt.Errorf("formatting mock file: %w", err)
	}
	return formatted, nil
}

func goimports(src []byte) ([]byte, error) {
	formatted, err := imports.Process("/", src, &imports.Options{
		TabWidth:   8,
		TabIndent:  true,
		Comments:   true,
		Fragment:   true,
		FormatOnly: true,
	})
	if err != nil {
		return nil, fmt.Errorf("goimports: %s", err)
	}

	return formatted, nil
}

func gofmt(src []byte) ([]byte, error) {
	formatted, err := format.Source(src)
	if err != nil {
		return nil, fmt.Errorf("go/format: %s", err)
	}

	return formatted, nil
}
