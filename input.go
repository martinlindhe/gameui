// parts from ebiten/examples/blocks/blocks/input.go

package ui

import (
	"github.com/hajimehoshi/ebiten"
)

// Input ...
type Input struct {
	keyStates   [256]int
	mouseStates [3]int
	X, Y        int
}

// StateForKey returns true once when the key starts to be pressed
func (i *Input) StateForKey(key Key) bool {
	return i.keyStates[key] == 1
}

// WasPressed returns true if key was pressed (for ctrl, alt etc)
func (i *Input) WasPressed(key Key) bool {
	return i.keyStates[key] != 0
}

// NewStateForMouse returns true if state was triggered
func (i *Input) NewStateForMouse(mouse MouseButton) bool {
	return i.mouseStates[mouse] == 1
}

// ContinuedStateForMouse returns true if state is not false
func (i *Input) ContinuedStateForMouse(mouse MouseButton) bool {
	return i.mouseStates[mouse] != 0
}

// ConsumeStateForMouse makes the mouse state been consumed, so click-thru doesn't happen
func (i *Input) ConsumeStateForMouse(mouse MouseButton) {
	i.mouseStates[mouse] = 2
}

// updateKeyboard ...
func (i *Input) updateKeyboard() {
	for key := range i.keyStates {
		if !ebiten.IsKeyPressed(ebiten.Key(key)) {
			i.keyStates[key] = 0
			continue
		}
		i.keyStates[key]++
	}
}

// updateMouse ...
func (i *Input) updateMouse() {
	for mouse := range i.mouseStates {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButton(mouse)) {
			i.mouseStates[mouse] = 0
			continue
		}
		i.mouseStates[mouse]++
	}
	i.X, i.Y = ebiten.CursorPosition()
}
