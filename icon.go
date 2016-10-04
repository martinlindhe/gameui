package ui

import "image"

// Icon ...
type Icon struct {
	component
	Image *image.RGBA
}

// NewIcon ...
func NewIcon(image *image.RGBA) *Icon {
	ico := Icon{Image: image}
	return &ico
}

// Draw ...
func (ico *Icon) Draw(mx, my int) *image.RGBA {
	return ico.Image
}
