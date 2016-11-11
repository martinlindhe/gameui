package ui

import (
	"image"
	"image/color"
	"io"
	"os"
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
func DrawRect(img *image.RGBA, r image.Rectangle, col color.Color) {
	vLine(img, r.Min.X, r.Min.Y, r.Max.Y, col) // left
	vLine(img, r.Max.X, r.Min.Y, r.Max.Y, col) // right

	hLine(img, r.Min.X+1, r.Max.X-1, r.Min.Y, col) // top
	hLine(img, r.Min.X+1, r.Max.X-1, r.Max.Y, col) // bottom
}

// OpenImage loads an image from file. based on Open from disintegration/imaging
// https://github.com/disintegration/imaging/blob/master/helpers.go#L68
func OpenImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return decodeImage(file)
}

// decodeImage reads an image from r
func decodeImage(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}
