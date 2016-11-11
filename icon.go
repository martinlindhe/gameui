package ui

import (
	"image"
	"image/draw"
)

// Icon is a icon (UI component)
type Icon struct {
	component
	icon *image.RGBA
}

// NewIcon creates a new icon
func NewIcon(img image.Image) *Icon {
	ico := Icon{}
	ico.SetIcon(img)
	b := ico.icon.Bounds()
	ico.Dimension.Width = b.Max.X
	ico.Dimension.Height = b.Max.Y
	return &ico
}

// SetIcon sets the icon
func (ico *Icon) SetIcon(img image.Image) {
	b := img.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, b.Min, draw.Src)
	ico.icon = m
}

// Draw ...
func (ico *Icon) Draw(mx, my int) *image.RGBA {
	if ico.isHidden {
		return nil
	}
	return ico.icon
}
