package ui

import (
	"image"
	"image/color"
)

// HLine draws a horizontal line
func HLine(img *image.RGBA, x1, x2, y int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(img *image.RGBA, x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// DrawRect ... XXX take image.Rectangle param
func DrawRect(img *image.RGBA, x0, y0, x1, y1 int, col color.Color) {
	// left, right
	VLine(img, x0, y0+1, y1, col)
	VLine(img, x1-1, y0+1, y1, col)
	// top, bottom
	HLine(img, x0+1, x1-1, y0, col)
	HLine(img, x0+1, x1-1, y1, col)
}
