package ui

import (
	"image"
)

// Component represents any type of UI component
type Component interface {
	Draw() (*image.RGBA, error)
	GetUpperLeft() (int, int)
}

// component is the abstract base class for ui components
type component struct {
	IsMouseOver   bool
	IsClean       bool // does component need redraw?
	Width, Height int
	X, Y          int
	Image         *image.RGBA
}

func (c component) GetUpperLeft() (int, int) { // XXX replace with GetBounds()
	return c.X, c.Y
}

// updateHover toggles IsMouseOver if cursor is over element
func (c component) updateHover(mx, my int) {
	c.IsMouseOver = mx >= c.X && mx <= c.X+c.Width &&
		my >= c.Y && my <= c.Y+c.Height
}
