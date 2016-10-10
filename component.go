package ui

import (
	"image"
	"image/draw"
	"log"
)

// Component represents any type of UI component
type Component interface {
	// Draw return nil if no image is drawn
	Draw(mx, my int) *image.RGBA
	GetBounds() (int, int, int, int)
	GetRect() image.Rectangle
	Hover(bool)
	IsClean() bool
	Click(Point)
}

// component is the abstract base class for ui components
type component struct {
	IsMouseOver   bool
	isClean       bool // does component need redraw?
	Width, Height int  // size of component
	Position      Point
	Image         *image.RGBA
	OnClick       func()
	children      []Component
}

func (c *component) Click(mouse Point) {
	if c.OnClick == nil {
		log.Println("error: OnClick == nil")
		return
	}
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

func (c *component) isChildrenClean() bool {
	for _, child := range c.children {
		if !child.IsClean() {
			return false
		}
	}
	return true
}

func (c *component) drawChildren(mx, my int) {
	for _, child := range c.children {
		img := child.Draw(mx, my)
		if img == nil {
			continue
		}
		x, y, w, h := child.GetBounds()
		x1 := x + w
		y1 := y + h
		child.Hover(mx >= x && mx <= x1 && my >= y && my <= y1)

		dr := image.Rect(x, y, x1, y1)
		draw.Draw(c.Image, dr, img, image.ZP, draw.Over)
	}
}
