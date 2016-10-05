package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// IconGroup is a tile-hased grid display of object icons
type IconGroup struct {
	component
	columns, rows int
}

// NewIconGroup ...
func NewIconGroup(columns, rows, width, height int) *IconGroup {
	igrp := IconGroup{}
	igrp.columns = columns
	igrp.rows = rows
	igrp.Width = width
	igrp.Height = height
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

	// XXX draw content

	grp.isClean = true
	return grp.Image
}
