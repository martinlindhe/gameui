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
	font  *Font
}

// NewText ...
func NewText(size float64, color color.Color) *Text {
	if size < 3 {
		log.Fatal("txt.size too small:", size)
	}

	txt := &Text{}
	txt.size = size
	txt.color = color

	var err error
	txt.font, err = NewFont(defaultFontName, txt.size, dpi, txt.color)
	if err != nil {
		log.Println("NewFont err", err)
		return nil
	}
	return txt
}

// SetText ...
func (txt *Text) SetText(s string) {
	txt.text = s
	txt.isClean = false
}

// GetWidth returns the rendered width in pixel
func (txt *Text) GetWidth() int {
	if txt.Image == nil {
		txt.Draw(-1, -1)
	}
	return txt.Width
}

// Draw redraws internal buffer
func (txt *Text) Draw(mx, my int) *image.RGBA {
	if txt.isClean {
		return txt.Image
	}

	img, err := txt.font.Print(txt.text)
	if err != nil {
		log.Println("Print err", err)
		return nil
	}

	b := img.Bounds()
	txt.Width = b.Max.X
	txt.Height = b.Max.Y
	txt.Image = img
	txt.isClean = true
	return img
}
