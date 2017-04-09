package ui

import (
	"image"
	"log"
)

// Text is a line of text (UI component)
type Text struct {
	component
	text string
	font *Font
}

// NewText creates a new Text instance
func NewText(font *Font) *Text {
	txt := Text{}
	txt.backgroundColor = Transparent
	txt.font = font
	txt.isClean = true
	return &txt
}

// SetFont sets the font to use
func (txt *Text) SetFont(font *Font) {
	txt.font = font
	txt.isClean = false
}

// SetText sets the text
func (txt *Text) SetText(s string) *Text {
	if s == txt.text {
		return txt
	}

	txt.isClean = false
	txt.text = s
	txt.isHidden = txt.text == ""

	img, err := txt.font.Print(txt.text)
	if err != nil {
		log.Println("Print err", err)
		return nil
	}

	txt.Image = img
	if img != nil {
		b := img.Bounds()
		txt.Dimension.Width = b.Max.X
		txt.Dimension.Height = b.Max.Y
	}
	return txt
}

// GetText returns the text
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
		txt.isClean = true
		return nil
	}
	return txt.Image
}
