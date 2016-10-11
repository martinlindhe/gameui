package ui

// XXX tests!!!

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// IconGroup is a tile-hased grid display of object icons
type IconGroup struct {
	component
	columns, rows         int
	iconWidth, iconHeight int // size of each icon
	objects               []IconGroupObject
}

// IconGroupObject is something that is contained in the icon group
type IconGroupObject interface {
	Name() string
	Icon() *image.RGBA
	Click()
	ID() uint64
}

const (
	iconBorderPad = 1
)

// NewIconGroup ...
func NewIconGroup(columns, rows, iconWidth, iconHeight int) *IconGroup {

	pad := 2 // 1 px border, + 1 px cell padding
	componentWidth := (columns * iconWidth) + (pad * 2)
	componentHeight := (rows * iconHeight) + (pad * 2)
	igrp := IconGroup{}
	igrp.columns = columns
	igrp.rows = rows
	igrp.Dimension.Width = componentWidth
	igrp.Dimension.Height = componentHeight
	igrp.iconWidth = iconWidth
	igrp.iconHeight = iconHeight
	return &igrp
}

// Draw redraws internal buffer
func (grp *IconGroup) Draw(mx, my int) *image.RGBA {
	if grp.Hidden {
		return nil
	}
	if grp.isClean {
		return grp.Image
	}

	rect := image.Rect(0, 0, grp.Dimension.Width, grp.Dimension.Height)
	if grp.Image == nil {
		grp.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(grp.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}

	// draw outline
	DrawRect(grp.Image, &rect, color.White)

	grp.drawIcons(mx, my)

	grp.isClean = true
	return grp.Image
}

// AddObject adds an object to display in the group
func (grp *IconGroup) AddObject(o IconGroupObject) {
	grp.objects = append(grp.objects, o)
	grp.isClean = false
}

// RemoveObjectByID ...
func (grp *IconGroup) RemoveObjectByID(id uint64) {
	for i, c := range grp.objects {
		if c.ID() == id {
			grp.objects = append(grp.objects[:i], grp.objects[i+1:]...)
			grp.isClean = false
			return
		}
	}
}

// ClearObjects removes all displayed content
func (grp *IconGroup) ClearObjects() {
	grp.objects = nil
	grp.isClean = false
}

func (grp *IconGroup) drawIcons(mx, my int) {
	x := iconBorderPad + 1
	y := iconBorderPad + 1
	col := 0
	row := 0

	// TODO: allow scrolling with scroll bar / mouse wheel

	for _, o := range grp.objects {
		// TODO show o.Name() if mouse is hover without click
		//fmt.Println("drawing", o.Name())

		img := o.Icon()
		if img == nil {
			fmt.Println("ERROR: UI IconGroup object", o.Name(), "lacks icon")
			continue
		}
		b := img.Bounds()
		w := b.Max.X
		h := b.Max.Y
		x1 := x + w
		y1 := y + h

		dr := image.Rect(x, y, x1, y1)
		draw.Draw(grp.Image, dr, img, image.ZP, draw.Over)
		x += w
		col++
		if col >= grp.columns {
			col = 0
			x = iconBorderPad + 1
			y += h
			row++
		}
		if row >= grp.rows {
			break
		}
	}
}

// Click pass click to child icon
func (grp *IconGroup) Click(mouse Point) {

	x := iconBorderPad + 1
	y := iconBorderPad + 1
	col := 0
	row := 0

	childPoint := Point{X: mouse.X - grp.Position.X, Y: mouse.Y - grp.Position.Y}

	for _, c := range grp.objects {
		b := c.Icon().Bounds()
		x1 := x + b.Max.X
		y1 := y + b.Max.Y
		r := image.Rect(x, y, x1, y1)
		if childPoint.In(r) {
			c.Click()
			return
		}

		x += b.Max.X
		col++
		if col >= grp.columns {
			col = 0
			x = iconBorderPad + 1
			y += b.Max.Y
			row++
		}
	}
}
