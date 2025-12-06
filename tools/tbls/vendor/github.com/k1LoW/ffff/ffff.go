package ffff

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/beta/freetype/truetype"
	"github.com/k1LoW/fontdir"
	"github.com/sahilm/fuzzy"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/text/unicode/norm"
)

type Font struct {
	Path string
	Face font.Face
}

var (
	names = []string{}
	paths = []string{}
	fonts = map[string]Font{}
)

// FuzzyFind find font by keyword
func FuzzyFind(keyword string, to *truetype.Options, oo *opentype.FaceOptions) (Font, error) {
	list := []string{}
	pathOnly := false
	lk := strings.ToLower(keyword)
	if strings.HasSuffix(lk, ".ttf") || strings.HasSuffix(lk, ".ttc") || strings.HasSuffix(lk, ".otf") {
		pathOnly = true
	}

	if len(fonts) == 0 {
		if err := listFonts(to, oo); err != nil {
			return Font{}, err
		}
	}

	if !pathOnly {
		list = append(list, names...)
	}
	list = append(list, paths...)

	matches := fuzzy.Find(keyword, list)
	if len(matches) == 0 {
		return Font{}, fmt.Errorf("could not find font: %s", keyword)
	}
	m, ok := fonts[matches[0].Str]
	if !ok {
		return Font{}, fmt.Errorf("could not find font: %s", keyword)
	}
	return m, nil
}

// FuzzyFindPath find font file path by keyword
func FuzzyFindPath(keyword string) (string, error) {
	f, err := FuzzyFind(keyword, nil, nil)
	if err != nil {
		return "", err
	}
	return f.Path, nil
}

// FuzzyFindFace find font.Face by keyword
func FuzzyFindFace(keyword string, to *truetype.Options, oo *opentype.FaceOptions) (font.Face, error) {
	f, err := FuzzyFind(keyword, to, oo)
	if err != nil {
		return nil, err
	}
	return f.Face, nil
}

func listFonts(to *truetype.Options, oo *opentype.FaceOptions) error {
	for _, dir := range fontdir.Get() {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			lp := strings.ToLower(path)
			abs, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			var face font.Face
			if strings.HasSuffix(lp, ".ttf") || strings.HasSuffix(lp, ".ttc") {
				// TrueType
				d, err := ioutil.ReadFile(filepath.Clean(path))
				if err != nil {
					return err
				}
				f, err := truetype.Parse(d)
				if err != nil {
					return nil
				}
				name := f.Name(4)
				names = append(names, name)
				face = truetype.NewFace(f, to)
				fonts[name] = Font{
					Path: abs,
					Face: face,
				}
			} else if strings.HasSuffix(lp, ".otf") {
				// OpenType
				d, err := ioutil.ReadFile(filepath.Clean(path))
				if err != nil {
					return err
				}
				f, err := sfnt.Parse(d)
				if err != nil {
					return nil
				}
				name, err := f.Name(nil, 4)
				if err != nil {
					return nil
				}
				names = append(names, name)
				face, err = opentype.NewFace(f, oo)
				if err != nil {
					return nil
				}
				fonts[name] = Font{
					Path: abs,
					Face: face,
				}
			} else {
				return nil
			}
			filename := norm.NFKC.String(filepath.Base(abs))
			paths = append(paths, filename)
			fonts[filename] = Font{
				Path: abs,
				Face: face,
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
