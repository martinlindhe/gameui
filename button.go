package ui

import (
	"image"
	"image/color"
	"image/draw"
	"log"

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

// SetImage a image to show on button, instead of text
func (btn *Button) SetImage(img *image.RGBA) {
	btn.image = img
	btn.IsClean = false
}

// XXX set text + render it

// Draw redraws internal buffer
func (btn *Button) Draw(mx, my int) *image.RGBA {

	// dont draw if nothing's changed
	if btn.IsClean {
		return btn.Image
	}

	rect := image.Rect(0, 0, btn.Width, btn.Height)
	im := image.NewRGBA(rect)

	// draw button outline
	game.DrawRect(im, &rect, color.White)

	if btn.image != nil {
		btn.drawImage(im)
	}
	btn.Image = im
	btn.IsClean = true
	return im
}

func (btn *Button) drawImage(im *image.RGBA) {
	allB := im.Bounds()
	btnB := btn.image.Bounds()
	if allB.Max.X > btn.Width || allB.Max.Y > btn.Height {
		log.Println("button.drawImage: image is bigger than container button")
	}

	// centered
	x0 := (allB.Size().X / 2) - (btnB.Size().X / 2)
	y0 := (allB.Size().Y / 2) - (btnB.Size().Y / 2)
	x1 := x0 + btnB.Max.X
	y1 := y0 + btnB.Max.Y

	rect := image.Rect(x0, y0, x1, y1)
	draw.Draw(im, rect, btn.image, allB.Min, draw.Over)
}
