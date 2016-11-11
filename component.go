package ui

import (
	"image"
	"image/color"
	"image/draw"
	"log"
)

// Component represents any type of UI component
type Component interface {
	Draw(mx, my int) *image.RGBA // return nil if no image is drawn
	GetBounds() image.Rectangle
	Hover(bool)
	IsClean() bool
	IsHidden() bool
	Click(Point) bool // return true if click was handled
}

// component is the abstract base class for ui components (doesn't implement Draw())
type component struct {
	IsMouseOver bool
	isClean     bool // does component need redraw?
	isHidden    bool
	Dimension   Dimension
	Position    Point
	Image       *image.RGBA
	OnClick     func()
	children    []Component
}

func (c *component) Click(mouse Point) bool {
	if c.OnClick == nil {
		log.Println("OnClick == nil for clicked component")
		return false
	}
	c.OnClick()
	return true
}

func (c *component) IsClean() bool {
	return c.isClean
}

func (c *component) IsHidden() bool {
	return c.isHidden
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

// RemoveAllChildren removes all children
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

func (c *component) Hide() {
	c.isHidden = true
	c.isClean = false
}

func (c *component) Show() {
	c.isHidden = false
	c.isClean = false
}

// Positioner ...
type Positioner interface {
	GetPosition() Point
	SetPosition(Point)
}

func (c *component) GetPosition() Point {
	return c.Position
}

func (c *component) SetPosition(pos Point) {
	c.Position = pos
}

func (c *component) initImage() {
	rect := image.Rect(0, 0, c.Dimension.Width, c.Dimension.Height)
	if c.Image == nil {
		c.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(c.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}
}
