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
	GetBounds() image.Rectangle
	Hover(bool)
	IsClean() bool
	Click(Point)
}

// component is the abstract base class for ui components (doesn't implement Draw())
type component struct {
	IsMouseOver bool
	isClean     bool // does component need redraw?
	Hidden      bool
	Dimension   Dimension
	Position    Point
	Image       *image.RGBA
	OnClick     func()
	children    []Component
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

func (c component) GetBounds() image.Rectangle {
	min := image.Point{c.Position.X, c.Position.Y}
	max := image.Point{c.Position.X + c.Dimension.Width, c.Position.Y + c.Dimension.Height}
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

func (c *component) RemoveAllChildren() {
	c.children = nil
}

func (c *component) drawChildren(mx, my int) {
	for _, child := range c.children {
		img := child.Draw(mx, my)
		if img == nil {
			continue
		}
		r := child.GetBounds()
		child.Hover(mx >= r.Min.X && mx <= r.Max.X && my >= r.Min.Y && my <= r.Max.Y)

		draw.Draw(c.Image, r, img, image.ZP, draw.Over)
	}
}
