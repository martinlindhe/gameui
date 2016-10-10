package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Window ...
type Window struct {
	component
	title    string
	children []Component
}

// NewWindow ...
func NewWindow(width, height int) *Window {
	wnd := Window{title: "new window"}
	wnd.Width = width
	wnd.Height = height
	return &wnd
}

// SetTitle ...
func (wnd *Window) SetTitle(s string) *Window {
	wnd.title = s
	return wnd
}

// AddChild ...
func (wnd *Window) AddChild(c Component) {
	wnd.children = append(wnd.children, c)
}

// Draw redraws internal buffer
func (wnd *Window) Draw(mx, my int) *image.RGBA {
	if wnd.isClean {
		return wnd.Image
	}

	rect := image.Rect(0, 0, wnd.Width, wnd.Height)
	if wnd.Image == nil {
		wnd.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(wnd.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}

	// draw outline
	DrawRect(wnd.Image, &rect, color.White)

	// XXX draw children
	for _, c := range wnd.children {
		img := c.Draw(mx, my)
		if img == nil {
			continue
		}
		x, y, w, h := c.GetBounds()
		x1 := x + w
		y1 := y + h
		c.Hover(mx >= x && mx <= x1 && my >= y && my <= y1)

		dr := image.Rect(x, y, x1, y1)
		draw.Draw(wnd.Image, dr, img, image.ZP, draw.Over)
	}

	wnd.isClean = true
	return wnd.Image
}
