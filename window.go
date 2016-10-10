package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Window ...
type Window struct {
	component
	title string
}

// NewWindow ...
func NewWindow(width, height int) *Window {
	wnd := Window{title: "new window"}
	wnd.Width = width
	wnd.Height = height
	return &wnd
}

// AddChild ...
func (wnd *Window) AddChild(c Component) {
	wnd.children = append(wnd.children, c)
}

// SetTitle ...
func (wnd *Window) SetTitle(s string) *Window {
	wnd.title = s
	return wnd
}

// Draw redraws internal buffer
func (wnd *Window) Draw(mx, my int) *image.RGBA {
	if wnd.isClean {
		if !wnd.isChildrenClean() {
			wnd.isClean = false
		} else {
			return wnd.Image
		}
	}

	rect := image.Rect(0, 0, wnd.Width, wnd.Height)
	if wnd.Image == nil {
		wnd.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(wnd.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}

	// draw outline
	DrawRect(wnd.Image, &rect, color.White)

	wnd.drawChildren(mx, my)

	wnd.isClean = true
	return wnd.Image
}

// Click pass click to window child components
func (wnd *Window) Click(mouse Point) {
	childPoint := Point{X: mouse.X - wnd.Position.X, Y: mouse.Y - wnd.Position.Y}
	for _, c := range wnd.children {
		if childPoint.In(c.GetRect()) {
			c.Click(childPoint)
			return
		}
	}
}
