package ui

import (
	"image"
	"image/color"
	"log"
)

const (
	dpi = 72
)

// Text ...
type Text struct {
	component
	text string
	size float64
}

// NewText ...
func NewText(text string, size float64) *Text {

	txt := &Text{}
	txt.SetText(text)
	txt.size = size
	return txt
}

// SetText ...
func (txt *Text) SetText(s string) {
	txt.text = s
	txt.IsClean = false
}

// Draw redraws internal buffer
func (txt *Text) Draw() *image.RGBA {

	// dont draw if nothing's changed
	if txt.IsClean {
		return txt.Image
	}
	if txt.size == 0 {
		panic("txt.size == 0")
	}
	// XXX use font every time. later, dont call NewFont so often!
	fnt, err := NewFont("../_resources/font/tiny/tiny.ttf", txt.size, dpi, color.White)
	if err != nil {
		log.Fatal(err)
	}

	img, err := fnt.Print(txt.text)
	if err != nil {
		log.Fatal(err)
	}

	b := img.Bounds()
	txt.Width = b.Max.X
	txt.Height = b.Max.Y
	txt.Image = img
	txt.IsClean = true
	return img
}
