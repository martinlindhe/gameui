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
	txt.text = text
	txt.size = size
	return txt
}

// SetText ...
func (txt *Text) SetText(s string) {
	txt.text = s
	txt.IsClean = false
}

// Draw redraws internal buffer
func (txt *Text) Draw() (*image.RGBA, error) {

	// dont draw if nothing's changed
	if txt.IsClean {
		return txt.Image, nil
	}

	// XXX use font every time. later, dont call NewText so often!
	fnt, err := NewFont("../_resources/font/tiny/tiny.ttf", txt.size, dpi, color.White)
	if err != nil {
		log.Fatal(err)
	}

	// XXX Print should only return as big image as needed. height is wrong!
	txt.Image, err = fnt.Print("HEJ")
	if err != nil {
		log.Fatal(err)
	}

	txt.IsClean = true
	return txt.Image, nil
}
