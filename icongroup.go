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
}

// NewIconGroup ...
func NewIconGroup(columns, rows, iconWidth, iconHeight int) *IconGroup {

	pad := 2 // 1 px border, + 1 px cell padding
	componentWidth := columns*iconWidth + pad*2
	componentHeight := rows*iconHeight + pad*2

	igrp := IconGroup{}
	igrp.columns = columns
	igrp.rows = rows
	igrp.Width = componentWidth
	igrp.Height = componentHeight
	igrp.iconWidth = iconWidth
	igrp.iconHeight = iconHeight
	return &igrp
}

// Draw redraws internal buffer
func (grp *IconGroup) Draw(mx, my int) *image.RGBA {
	if grp.isClean {
		return grp.Image
	}

	rect := image.Rect(0, 0, grp.Width, grp.Height)
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

func (grp *IconGroup) drawIcons(mx, my int) {
	fmt.Println("drawIcons")
	pad := 1
	x := pad + 1
	y := pad + 1
	col := 0

	for _, o := range grp.objects {
		// TODO show o.Name() if mouse is hover without click
		fmt.Println("drawing", o.Name())

		img := o.Icon()
		b := img.Bounds()
		w := b.Max.X
		h := b.Max.Y
		x1 := x + w
		y1 := y + h

		// XXX handle click within bounds
		// XXX child.Hover(mx >= x && mx <= x1 && my >= y && my <= y1)

		dr := image.Rect(x, y, x1, y1)
		draw.Draw(grp.Image, dr, img, image.ZP, draw.Over)
		x += w
		col++
		if col >= grp.columns {
			x = pad + 1
			y += h
		}
	}
}
