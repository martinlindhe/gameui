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
	txt.font, err = NewFont(defaultFontName, size, dpi, txt.color)
	if err != nil {
		log.Println("NewFont err", err)
		return nil
	}
	return txt
}

// SetText ...
func (txt *Text) SetText(s string) *Text {
	if s != txt.text {
		txt.isClean = false
	}
	txt.text = s
	return txt
}

// GetText ...
func (txt *Text) GetText() string {
	return txt.text
}

// GetWidth returns the rendered width in pixel
func (txt *Text) GetWidth() int {
	if txt.Image == nil {
		txt.Draw(-1, -1)
	}
	return txt.Dimension.Width
}

// Draw redraws internal buffer
func (txt *Text) Draw(mx, my int) *image.RGBA {
	if txt.Hidden {
		return nil
	}
	if txt.isClean {
		return txt.Image
	}

	img, err := txt.font.Print(txt.text)
	if err != nil {
		log.Println("Print err", err)
		return nil
	}

	b := img.Bounds()
	txt.Image = img
	txt.Dimension.Width = b.Max.X
	txt.Dimension.Height = b.Max.Y
	txt.isClean = true
	return img
}
