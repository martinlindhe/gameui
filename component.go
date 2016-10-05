package ui

import "image"

// Component represents any type of UI component
type Component interface {
	// Draw return nil if no image is drawn
	Draw(mx, my int) *image.RGBA
	GetBounds() (int, int, int, int)
	GetRect() image.Rectangle
	Hover(bool)
	IsClean() bool
	Click()
}

// component is the abstract base class for ui components
type component struct {
	IsMouseOver   bool
	isClean       bool // does component need redraw?
	Width, Height int
	Position      Point
	Image         *image.RGBA
	OnClick       func()
}

func (c *component) Click() {
	c.OnClick()
}

func (c *component) IsClean() bool {
	return c.isClean
}

// GetBounds returns x, y, width, height
func (c component) GetBounds() (int, int, int, int) {
	return c.Position.X, c.Position.Y, c.Width, c.Height
}

func (c component) GetRect() image.Rectangle {
	min := image.Point{c.Position.X, c.Position.Y}
	max := image.Point{c.Position.X + c.Width, c.Position.Y + c.Height}
	return image.Rectangle{min, max}
}

// set to true when mouse is hovering component
func (c component) Hover(b bool) {
	c.IsMouseOver = b
}
