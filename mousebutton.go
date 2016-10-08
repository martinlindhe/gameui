// copy from ebiten/mosuebutton.go

package ui

import "github.com/hajimehoshi/ebiten"

// A MouseButton represents a mouse button.
type MouseButton int

// MouseButtons
const (
	MouseButtonLeft   = MouseButton(ebiten.MouseButtonLeft)
	MouseButtonRight  = MouseButton(ebiten.MouseButtonRight)
	MouseButtonMiddle = MouseButton(ebiten.MouseButtonMiddle)
)
