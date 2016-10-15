package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Button ...
type Button struct {
	component
	icon *image.RGBA
	Text *Text
}

// NewButton ...
func NewButton(width, height int) *Button {
	btn := &Button{}
	btn.Dimension.Width = width
	btn.Dimension.Height = height
	btn.Text = NewText(float64(height-1), White)
	return btn
}

// SetIcon a image to show on button, instead of text
func (btn *Button) SetIcon(img *image.RGBA) {
	btn.icon = img
	btn.isClean = false
}

// SetText ...
func (btn *Button) SetText(s string) *Button {
	btn.Text.SetText(s)
	return btn
}

// Draw redraws internal buffer
func (btn *Button) Draw(mx, my int) *image.RGBA {
	if btn.Hidden {
		return nil
	}
	if btn.isClean {
		return btn.Image
	}

	rect := image.Rect(0, 0, btn.Dimension.Width, btn.Dimension.Height)
	if btn.Image == nil {
		btn.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(btn.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}

	// draw outline
	outlineRect := image.Rect(0, 0, btn.Dimension.Width-1, btn.Dimension.Height-1)
	DrawRect(btn.Image, &outlineRect, White)

	btn.drawIcon()

	if btn.Text.text != "" {
		txt := btn.Text.Draw(mx, my)
		// XXX center text ?
		b := txt.Bounds()
		x0 := 2
		y0 := 1
		x1 := x0 + b.Max.X
		y1 := y0 + b.Max.Y
		textRect := image.Rect(x0, y0, x1, y1)
		draw.Draw(btn.Image, textRect, txt, image.ZP, draw.Over)
	}

	btn.isClean = true
	return btn.Image
}

func (btn *Button) drawIcon() {
	if btn.icon == nil {
		return
	}
	allB := btn.Image.Bounds()
	btnB := btn.icon.Bounds()

	// centered
	x0 := (allB.Size().X / 2) - (btnB.Size().X / 2)
	y0 := (allB.Size().Y / 2) - (btnB.Size().Y / 2)
	x1 := x0 + btnB.Max.X
	y1 := y0 + btnB.Max.Y

	rect := image.Rect(x0, y0, x1, y1)
	draw.Draw(btn.Image, rect, btn.icon, image.ZP, draw.Over)
}
