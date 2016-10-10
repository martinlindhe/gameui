package ui

import "image"

// Point ...
type Point image.Point

// In ...
func (p *Point) In(rect image.Rectangle) bool {
	return image.Point(*p).In(rect)
}
