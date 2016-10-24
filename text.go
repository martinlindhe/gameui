package ui

import (
	"image"
	"image/color"
	"log"
)

// Text ...
type Text struct {
	component
	text  string
	size  float64
	dpi   float64
	color color.Color
	font  *Font
}

// NewText ...
func NewText(size float64, color color.Color) *Text {
	txt := Text{}
	txt.size = size
	txt.color = color
	txt.dpi = 72

	if err := txt.setFont(defaultFontName); err != nil {
		log.Println("NewText err", err)
		return nil
	}
	return &txt
}

func (txt *Text) setFont(fontName string) error {
	var err error
	txt.font, err = NewFont(fontName, txt.size, txt.dpi, txt.color)
	return err
}

// SetText ...
func (txt *Text) SetText(s string) *Text {
	if s != txt.text {
		txt.isClean = false
		txt.text = s
	}
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

// GetHeight returns the rendered height in pixel
func (txt *Text) GetHeight() int {
	if txt.Image == nil {
		txt.Draw(-1, -1)
	}
	return txt.Dimension.Height
}

// Draw redraws internal buffer
func (txt *Text) Draw(mx, my int) *image.RGBA {
	if txt.isHidden {
		return nil
	}
	if txt.isClean {
		return txt.Image
	}
	if txt.text == "" {
		return nil
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
