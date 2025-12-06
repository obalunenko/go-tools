# ffff [![Build Status](https://github.com/k1LoW/ffff/workflows/build/badge.svg)](https://github.com/k1LoW/ffff/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/k1LoW/ffff.svg)](https://pkg.go.dev/github.com/k1LoW/ffff)

**f**inder **f**or **f**ont **f**iles

## Usage

``` go
package main

import (
	"fmt"

	"github.com/beta/freetype/truetype"
	"github.com/k1LoW/ffff"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func main() {
	fontSize := 12
	dpi := 72
	to := &truetype.Options{
		Size:              fontSize,
		DPI:               dpi,
		Hinting:           font.HintingNone,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}
	oo := &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingNone,
	}

	face, err := ffff.FuzzyFindFace("Arial", to, oo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", face)
}
```
