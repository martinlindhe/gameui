package ui

import (
	"image"
	"image/color"
	"log"
)

const (
	dpi = 72
)

var (
	defaultFontName = assetPath("_resources/font/tiny/tiny.ttf")
)

// Text ...
type Text struct {
	component
	text  string
	size  float64
	color color.Color
}

// NewText ...
func NewText(text string, size float64, color color.Color) *Text {

	txt := &Text{}
	txt.SetText(text)
	txt.size = size
	txt.color = color
	return txt
}

// SetText ...
func (txt *Text) SetText(s string) {
	txt.text = s
	txt.IsClean = false
}

// Draw redraws internal buffer
func (txt *Text) Draw(mx, my int) *image.RGBA {

	// dont draw if nothing's changed
	if txt.IsClean {
		return txt.Image
	}
	if txt.size == 0 {
		log.Fatal("txt.size == 0")
	}
	// XXX use font every time. later, dont call NewFont so often!
	fnt, err := NewFont(defaultFontName, txt.size, dpi, txt.color)
	if err != nil {
		log.Println("NewFont err", err)
		return nil
	}

	img, err := fnt.Print(txt.text)
	if err != nil {
		log.Println("Print err", err)
		return nil
	}

	b := img.Bounds()
	txt.Width = b.Max.X
	txt.Height = b.Max.Y
	txt.Image = img
	txt.IsClean = true
	return img
}
