package ui

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	defaultFontName = assetPath("_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf")
	tinyFontName    = assetPath("_resources/font/tiny/tiny.ttf")
)

// Font is a font resource
type Font struct {
	dpi          float64
	size         float64
	spacing      float64
	font         *truetype.Font
	drawer       *font.Drawer
	color        color.Color
	cachedPrints map[string]*image.RGBA
}

const (
	fontRenderCache = 10
)

// NewFont prepares a new font resource for use
func NewFont(fontName string, size float64, dpi float64, col color.Color) (*Font, error) {
	b, err := ioutil.ReadFile(fontName)
	if err != nil {
		return nil, err
	}

	fnt := Font{}
	fnt.font, err = truetype.Parse(b)
	fnt.size = size
	fnt.dpi = dpi
	fnt.spacing = 1
	fnt.color = col
	fnt.cachedPrints = make(map[string]*image.RGBA)
	fnt.drawer = &font.Drawer{
		Src: image.NewUniform(col),
		Face: truetype.NewFace(fnt.font, &truetype.Options{
			Size:    fnt.size,
			DPI:     fnt.dpi,
			Hinting: font.HintingFull,
		}),
	}
	return &fnt, err
}

// Print draws text using the font
func (fnt *Font) Print(text string) (*image.RGBA, error) {
	if text == "" {
		fmt.Println("ERROR: font.Print with no text")
	}

	if val, ok := fnt.cachedPrints[text]; ok {
		return val, nil
	}

	if fnt.size == 0 {
		panic("fnt.size == 0")
	}

	dim := fnt.findDimension(text)
	fnt.drawer.Dst = image.NewRGBA(image.Rect(0, 0, dim.Width, dim.Height))

	dy := (fnt.size * fnt.dpi) / 72
	fnt.drawer.Dot = fixed.P(0, int(dy))
	fnt.drawer.DrawString(text)

	if img, ok := fnt.drawer.Dst.(*image.RGBA); ok {
		if len(fnt.cachedPrints) >= fontRenderCache {
			// trim cache. keep last few rendered strings
			randKey := getRandomKey(fnt.cachedPrints)
			delete(fnt.cachedPrints, randKey)
		}

		fnt.cachedPrints[text] = img
		return img, nil
	}
	return nil, fmt.Errorf("unhandled type %#v", fnt.drawer.Dst)
}

// Measure the text to calculate the minimum size of the image
func (fnt *Font) findDimension(text string) (dim Dimension) {
	dim.Width = int(fnt.drawer.MeasureString(text).Ceil())
	dim.Height = 4 + int(math.Ceil(fnt.size*fnt.dpi/72))
	return
}

func getRandomKey(m map[string]*image.RGBA) string {
	for k := range m {
		return k
	}
	return ""
}
