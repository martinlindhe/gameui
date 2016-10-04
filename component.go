package ui

import "image"

// Component represents any type of UI component
type Component interface {
	Draw(mx, my int) *image.RGBA
	SetParent(ui *UI)
	GetBounds() (int, int, int, int)
	Hover(bool)
}

// component is the abstract base class for ui components
type component struct {
	IsMouseOver   bool
	IsClean       bool // does component need redraw?
	Width, Height int
	Position      image.Point
	Image         *image.RGBA
	parent        *UI
}

func (c *component) SetParent(ui *UI) {
	c.parent = ui
}

// GetBounds returns x, y, width, height
func (c component) GetBounds() (int, int, int, int) {
	return c.Position.X, c.Position.Y, c.Width, c.Height
}

// set to true when mouse is hovering component
func (c component) Hover(b bool) {
	c.IsMouseOver = b
}
