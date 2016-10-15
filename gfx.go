package ui

import (
	"image"
	"image/color"
)

// hLine draws a horizontal line
func hLine(img *image.RGBA, x1, x2, y int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// vLine draws a veritcal line
func vLine(img *image.RGBA, x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// DrawRect ...
func DrawRect(img *image.RGBA, r *image.Rectangle, col color.Color) {
	// left, right
	vLine(img, r.Min.X, r.Min.Y+1, r.Max.Y-1, col)
	vLine(img, r.Max.X-1, r.Min.Y+1, r.Max.Y-1, col)
	// top, bottom
	hLine(img, r.Min.X, r.Max.X-1, r.Min.Y, col)
	hLine(img, r.Min.X, r.Max.X-1, r.Max.Y-1, col)
}
