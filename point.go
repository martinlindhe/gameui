package ui

import "image"

// Point is a absolute position
type Point image.Point

// In returns true if `p` is inside of `rect`
func (p *Point) In(rect image.Rectangle) bool {
	return image.Point(*p).In(rect)
}
