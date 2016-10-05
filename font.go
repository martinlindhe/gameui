package ui

//  TODO check out https://github.com/hajimehoshi/ebiten/blob/master/examples/common/font.go

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// Font represents a font resource
type Font struct {
	dpi          float64
	size         float64
	font         *truetype.Font
	drawer       *font.Drawer
	cachedPrints map[string]*image.RGBA
}

const (
	fontRenderCache = 10
)

// NewFont prepares a new font resource for use
func NewFont(fontName string, size float64, dpi float64, col color.Color) (*Font, error) {
	f, err := os.Open(fontName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var fnt Font
	fnt.font, err = truetype.Parse(b)
	fnt.size = size
	fnt.dpi = dpi
	fnt.cachedPrints = make(map[string]*image.RGBA)

	fnt.drawer = &font.Drawer{
		Src: image.NewUniform(col),
		Face: truetype.NewFace(fnt.font, &truetype.Options{
			Size:    fnt.size,
			DPI:     fnt.dpi,
			Hinting: font.HintingFull,
		}),
	}

	// XXX create a one-row wide lookup image with all letters rendered
	return &fnt, err
}

// StringInPixels ...
func (fnt *Font) StringInPixels(s string) int {
	return int(fnt.drawer.MeasureString(s).Ceil())
}

// Print draws text using the font
func (fnt *Font) Print(text string) (*image.RGBA, error) {

	// keep cache of last 10 rendered strings
	if val, ok := fnt.cachedPrints[text]; ok {
		return val, nil
	}

	width := fnt.StringInPixels(text)
	if fnt.size == 0 {
		panic("fnt.size == 0")
	}
	height := int(fnt.size - 1) // XXX not perfect
	fnt.drawer.Dst = image.NewRGBA(image.Rect(0, 0, width, height))

	dy := (fnt.size * fnt.dpi) / 72
	fnt.drawer.Dot = fixed.P(0, int(dy-2))
	fnt.drawer.DrawString(text)

	// trim cache
	if len(fnt.cachedPrints) > fontRenderCache {
		randKey := getRandomKey(fnt.cachedPrints)
		delete(fnt.cachedPrints, randKey)
	}

	fnt.cachedPrints[text] = fnt.drawer.Dst.(*image.RGBA)

	if img, ok := fnt.drawer.Dst.(*image.RGBA); ok {
		return img, nil
	}
	return nil, fmt.Errorf("bad print")
}

func getRandomKey(m map[string]*image.RGBA) string {
	for k := range m {
		return k
	}
	return ""
}
