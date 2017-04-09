package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Component represents any type of UI component
type Component interface {
	Draw(mx, my int) *image.RGBA // returns nil if no image is drawn
	GetBounds() image.Rectangle
	Hover(bool)
	Move(int, int)
	IsClean() bool
	IsHidden() bool
	Click(Point) bool
	Drag(bool)
	IsMouseOver() bool
	IsDraggable() bool
	IsDragged() bool
	Tooltip() *Button
}

// component is the abstract base class for ui components (doesn't implement Draw())
type component struct {
	isMouseOver     bool
	isClean         bool
	isHidden        bool
	isDraggable     bool
	isDragged       bool
	Dimension       Dimension
	Position        Point
	Image           *image.RGBA
	OnClick         func()
	children        []Component
	tooltip         *Button
	backgroundColor color.Color
}

var (
	tooltipBgColor = color.RGBA{0x50, 0x50, 0x50, 192} // grey, 25% transparent
)

// AddChild adds a child component to the Group
func (c *component) addChild(child Component) {
	c.isClean = false
	c.children = append(c.children, child)
}

// Click returns true if click was handled
func (c *component) Click(mouse Point) bool {
	if c.OnClick == nil {
		return false
	}
	c.OnClick()
	return true
}

// SetDraggable toggles the ability to drag the component
func (c *component) SetDraggable(enabled bool) {
	c.isDraggable = enabled
}

// Drag toggles if component is being dragged
func (c *component) Drag(enabled bool) {
	c.isDragged = enabled
}

// IsClean returns false if component needs redraw
func (c *component) IsClean() bool {
	return c.isClean
}

func (c *component) IsHidden() bool {
	return c.isHidden
}

func (c *component) IsMouseOver() bool {
	return c.isMouseOver
}

func (c *component) IsDraggable() bool {
	return c.isDraggable
}

func (c *component) IsDragged() bool {
	return c.isDragged
}

func (c *component) GetBounds() image.Rectangle {
	min := image.Point{c.Position.X, c.Position.Y}
	max := image.Point{c.Position.X + c.Dimension.Width, c.Position.Y + c.Dimension.Height}
	return image.Rectangle{min, max}
}

// SetTooltip sets the tooltip
func (c *component) SetTooltip(s string) {
	if c.tooltip != nil && c.tooltip.Text.text == s {
		return
	}
	tinyFont, _ := NewFont(defaultFontName, 10, 72, White)

	dim := tinyFont.findDimension(s)
	btn := NewButton(dim.Width+4, dim.Height)
	btn.SetBorderColor(White)
	btn.SetBackgroundColor(tooltipBgColor)
	btn.SetText(tinyFont, s)
	c.tooltip = btn
}

func (c *component) Move(x, y int) {
	c.Position.X = x
	c.Position.Y = y
}

// Tooltip returns the current tooltip
func (c *component) Tooltip() *Button {
	return c.tooltip
}

// Hover sets the mouse hovering state for component
func (c *component) Hover(b bool) {
	c.isMouseOver = b
}

func (c *component) isChildrenClean() bool {
	for _, child := range c.children {
		if !child.IsClean() {
			return false
		}
	}
	return true
}

// RemoveAllChildren removes all children from the component
func (c *component) RemoveAllChildren() {
	c.children = nil
}

// mx, my is absolute mouse position
func (c *component) drawChildren(mx, my int) {
	for _, child := range c.children {
		r := child.GetBounds()
		img := child.Draw(mx, my)
		if img != nil {
			draw.Draw(c.Image, r, img, image.ZP, draw.Over)
		}
	}
}

// Hide makes component hidden
func (c *component) Hide() {
	c.isHidden = true
	c.isClean = false
}

// Show makes component visible
func (c *component) Show() {
	c.isHidden = false
	c.isClean = false
}

// SetVisibility changes visibility of component
func (c *component) SetVisibility(b bool) {
	c.isHidden = !b
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

// SetBackgroundColor sets the background color
func (c *component) SetBackgroundColor(col color.Color) {
	c.backgroundColor = col
}

func (c *component) initImage() {
	rect := image.Rect(0, 0, c.Dimension.Width, c.Dimension.Height)
	if c.Image == nil {
		c.Image = image.NewRGBA(rect)
	}
	draw.Draw(c.Image, rect, &image.Uniform{c.backgroundColor}, image.ZP, draw.Src)
}
