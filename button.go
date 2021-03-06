package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Button is a button (UI component)
type Button struct {
	component
	icon        image.Image
	Text        *Text
	borderColor color.Color
}

var (
	buttonBorderColor = color.RGBA{0x50, 0x50, 0x50, 192} // grey, 75% transparent
)

// NewButton ...
func NewButton(width, height int) *Button {
	btn := Button{}
	btn.backgroundColor = Transparent
	btn.borderColor = buttonBorderColor
	btn.Dimension.Width = width
	btn.Dimension.Height = height
	return &btn
}

// SetBorderColor sets the border color
func (btn *Button) SetBorderColor(c color.Color) {
	btn.borderColor = c
}

// SetIcon a image to show on button, instead of text
func (btn *Button) SetIcon(img image.Image) {
	btn.icon = img
	btn.isClean = false
}

// SetText sets the button text
func (btn *Button) SetText(fnt *Font, s string) *Button {
	if btn.Text != nil && s == btn.Text.text {
		return btn
	}
	btn.Text = NewText(fnt)
	btn.Text.SetText(s)
	return btn
}

// Draw redraws internal buffer
func (btn *Button) Draw(mx, my int) *image.RGBA {
	if btn.isHidden {
		btn.isClean = true
		return nil
	}
	if btn.isClean {
		return btn.Image
	}
	btn.initImage()

	// draw outline
	outlineRect := image.Rect(0, 0, btn.Dimension.Width-1, btn.Dimension.Height-1)
	DrawRect(btn.Image, outlineRect, btn.borderColor)

	btn.drawIcon()

	if btn.Text != nil && btn.Text.text != "" {
		txt := btn.Text.Draw(mx, my)

		// left-aligned
		b := txt.Bounds()
		x0 := 1
		y0 := 0
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
