package ui

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/martinlindhe/farm/game"
)

// Button ...
type Button struct {
	component
	image *image.RGBA
}

// NewButton ...
func NewButton(width, height int) *Button {
	btn := &Button{}
	btn.Width = width
	btn.Height = height
	return btn
}

// SetImage ...
func (btn *Button) SetImage(img *image.RGBA) {
	btn.image = img
}

// Draw ...
func (btn *Button) Draw() (*image.RGBA, error) {

	// XXX reuse internal image buffer
	rect := image.Rect(0, 0, btn.Width, btn.Height)
	im := image.NewRGBA(rect)

	// draw button outline
	game.DrawRect(im, &rect, color.White)

	if btn.image != nil {
		// XXX draw image centered in btn.
		// XXX warn if image is bigger than btn. then auto shrink?

		b := btn.image.Bounds()
		draw.Draw(im, b, btn.image, b.Min, draw.Src)
	}
	return im, nil
}
