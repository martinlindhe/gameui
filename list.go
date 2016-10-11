package ui

// XXX a list is a component holding a number of rows of text, each is clickable

import (
	"image"
	"image/color"
	"image/draw"
)

// List ...
type List struct {
	component
	lines []Line
}

// Line ...
type Line interface {
	Name() string
	Color() color.Color
}

// NewList ...
func NewList(width, height int) *List {
	lst := List{}
	y := 0
	rowHeight := 12 // XXX font height
	for i := 0; i < 20; i++ {
		h := NewText(12, color.White)
		h.SetText("XXX")
		h.Position = Point{X: 0, Y: y}
		h.Dimension = Dimension{Width: width, Height: rowHeight}
		lst.children = append(lst.children, h)
		y += rowHeight
	}
	return &lst
}

// AddLine ...
func (lst *List) AddLine(l Line) {
	lst.lines = append(lst.lines, l)
}

// Draw redraws internal buffer
func (lst *List) Draw(mx, my int) *image.RGBA {
	if lst.Hidden {
		return nil
	}
	if lst.isClean {
		if !lst.isChildrenClean() {
			lst.isClean = false
		} else {
			return lst.Image
		}
	}

	rect := image.Rect(0, 0, lst.Dimension.Width, lst.Dimension.Height)
	if lst.Image == nil {
		lst.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(lst.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}

	// draw background color
	// XXX draw.Draw(lst.Image, rect, &image.Uniform{lst.backgroundColor}, image.ZP, draw.Over)

	lst.drawChildren(mx, my)

	lst.isClean = true
	return lst.Image
}

// Click pass click to window child components
func (lst *List) Click(mouse Point) {
	childPoint := Point{X: mouse.X - lst.Position.X, Y: mouse.Y - lst.Position.Y}
	for _, c := range lst.children {
		if childPoint.In(c.GetBounds()) {
			c.Click(childPoint)
			return
		}
	}
}
