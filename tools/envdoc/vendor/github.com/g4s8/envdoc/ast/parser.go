package ast

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"

	"github.com/g4s8/envdoc/utils"
)

type ParserConfigOption func(*Parser)

func WithDebug(debug bool) ParserConfigOption {
	return func(p *Parser) {
		p.debug = debug
	}
}

func WithExecConfig(execFile string, execLine int) ParserConfigOption {
	return func(p *Parser) {
		p.gogenFile = execFile
		p.gogenLine = execLine
	}
}

type Parser struct {
	fileGlob  string
	typeGlob  string
	gogenLine int
	gogenFile string
	debug     bool
}

func NewParser(fileGlob, typeGlob string, opts ...ParserConfigOption) *Parser {
	p := &Parser{
		fileGlob: fileGlob,
		typeGlob: typeGlob,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *Parser) Parse(dir string) ([]*FileSpec, error) {
	fset := token.NewFileSet()

	var colOpts []RootCollectorOption
	if p.typeGlob == "" {
		colOpts = append(colOpts, WithGoGenDecl(p.gogenLine, p.gogenFile))
	} else {
		m, err := utils.NewGlobMatcher(p.typeGlob)
		if err != nil {
			return nil, fmt.Errorf("create type glob matcher: %w", err)
		}
		colOpts = append(colOpts, WithTypeGlob(m))
	}
	if p.fileGlob != "" {
		m, err := utils.NewGlobMatcher(p.fileGlob)
		if err != nil {
			return nil, fmt.Errorf("create file glob matcher: %w", err)
		}
		colOpts = append(colOpts, WithFileGlob(m))
	}

	col := NewRootCollector(dir, colOpts...)

	if p.debug {
		fmt.Printf("Parsing dir %q (f=%q t=%q)\n", dir, p.fileGlob, p.typeGlob)
	}
	// walk through the directory and each subdirectory and call parseDir for each of them
	if err := filepath.Walk(dir, parseWalker(fset, col)); err != nil {
		return nil, fmt.Errorf("failed to walk through dir: %w", err)
	}

	if p.debug {
		fmt.Printf("Parsed types:\n")
		printTraverse(col.Files(), 0)
	}

	return col.Files(), nil
}

func parseWalker(fset *token.FileSet, col *RootCollector) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk through dir: %w", err)
		}
		if !info.IsDir() {
			return nil
		}
		if err := parseDir(path, fset, col); err != nil {
			return fmt.Errorf("failed to parse dir %q: %w", path, err)
		}
		return nil
	}
}

func parseDir(dir string, fset *token.FileSet, col *RootCollector) error {
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		return fmt.Errorf("failed to parse dir: %w", err)
	}

	for _, pkg := range pkgs {
		Walk(pkg, fset, col)
	}
	return nil
}
