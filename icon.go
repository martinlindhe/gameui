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
	b := ico.Image.Bounds()
	ico.Dimension.Width = b.Max.X
	ico.Dimension.Height = b.Max.Y
	return &ico
}

// Draw ...
func (ico *Icon) Draw(mx, my int) *image.RGBA {
	if ico.isHidden {
		return nil
	}
	return ico.Image
}
