package ui

import "image/color"

// Color ...
type Color color.Color

// ...
var (
	Black       = color.Black
	White       = color.White
	Red         = color.RGBA{255, 0, 0, 255}
	Green       = color.RGBA{0, 255, 0, 255}
	Blue        = color.RGBA{0, 0, 255, 255}
	Yellow      = color.RGBA{255, 255, 0, 255}
	Purple      = color.RGBA{255, 0, 255, 255}
	Brown       = color.RGBA{139, 69, 19, 255}
	Orange      = color.RGBA{255, 165, 0, 255}
	Pink        = color.RGBA{255, 105, 180, 255}
	DarkGrey    = color.RGBA{169, 169, 169, 255}
	LightGrey   = color.RGBA{211, 211, 211, 255}
	Transparent = color.Transparent
	Opaque      = color.Opaque
)
