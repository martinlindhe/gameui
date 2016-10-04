package ui

import "image"

// Icon ...
type Icon struct {
	component
	Image *image.RGBA
}

// NewIcon ...
func NewIcon(image *image.RGBA) *Icon {
	ico := &Icon{}
	ico.Image = image
	return ico
}

// Draw redraws internal buffer
func (ico *Icon) Draw(mx, my int) *image.RGBA {
	return ico.Image
}
