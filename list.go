package ui

import (
	"image"
	"image/color"
)

// List holds a number of rows of text, each is clickable (UI component)
type List struct {
	component
	rowHeight int
}

// Line defines the interface for lines of text usable with the List object
type Line interface {
	Name() string
	Color() color.Color
}

// NewList ...
func NewList(width, height int) *List {
	lst := List{}
	lst.Dimension = Dimension{Width: width, Height: height}
	lst.rowHeight = 12 // XXX
	return &lst
}

// addChild ...
func (lst *List) addChild(c Component) {
	lst.children = append(lst.children, c)
	lst.isClean = false
}

// AddLine ...
func (lst *List) AddLine(l Line, fnc func()) {
	h := NewText(float64(lst.rowHeight), l.Color())
	h.OnClick = fnc
	h.SetText(l.Name())
	h.Position = Point{X: 0, Y: len(lst.children) * lst.rowHeight}
	h.Dimension = Dimension{Width: lst.Dimension.Width, Height: lst.rowHeight}
	lst.addChild(h)
	lst.isClean = false
}

// Draw redraws internal buffer
func (lst *List) Draw(mx, my int) *image.RGBA {
	if lst.isHidden {
		return nil
	}
	if lst.isClean {
		if !lst.isChildrenClean() {
			lst.isClean = false
		} else {
			return lst.Image
		}
	}
	lst.initImage()

	// draw background color
	lst.drawChildren(mx, my)

	lst.isClean = true
	return lst.Image
}

// Click pass click to window child components
func (lst *List) Click(mouse Point) bool {
	childPoint := Point{X: mouse.X - lst.Position.X, Y: mouse.Y - lst.Position.Y}
	for _, c := range lst.children {
		if childPoint.In(c.GetBounds()) {
			c.Click(childPoint)
			return true
		}
	}
	return false
}
