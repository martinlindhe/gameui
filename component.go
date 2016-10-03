package ui

import (
	"fmt"
	"image"
)

// Component represents any type of UI component
type Component interface {
	Draw() *image.RGBA
	GetBounds() (int, int, int, int)
}

// component is the abstract base class for ui components, implementing Component interface
type component struct {
	IsMouseOver   bool
	IsClean       bool // does component need redraw?
	Width, Height int
	Position      image.Point
	Image         *image.RGBA
}

func (c component) Draw() *image.RGBA {
	fmt.Println("STUB Draw() - bug: child component must implement me!")
	return nil
}

// GetBounds returns x, y, width, height
func (c component) GetBounds() (int, int, int, int) {
	return c.Position.X, c.Position.Y, c.Width, c.Height
}

// updateHover toggles IsMouseOver if cursor is over element
func (c component) updateHover(mx, my int) {
	c.IsMouseOver = mx >= c.Position.X && mx <= c.Position.X+c.Width &&
		my >= c.Position.Y && my <= c.Position.Y+c.Height
}
