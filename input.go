// parts from ebiten/examples/blocks/blocks/input.go

package ui

import "github.com/hajimehoshi/ebiten"

// Input ...
type Input struct {
	keyStates   [256]int
	mouseStates [3]int
	X, Y        int
}

// StateForKey ...
func (i *Input) StateForKey(key Key) bool {
	return i.keyStates[key] == 1
}

// StateForMouse ...
func (i *Input) StateForMouse(mouse MouseButton) bool {
	return i.mouseStates[mouse] == 1
}

// ConsumeStateForMouse makes the mouse state been consumed, so click-thru doesn't happen
func (i *Input) ConsumeStateForMouse(mouse MouseButton) {
	i.mouseStates[mouse] = 2
}

// ClearMouse clear the mouse state
func (i *Input) ClearMouse() {
	for mouse := range i.mouseStates {
		i.mouseStates[mouse] = 0
	}
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
	mx, my := ebiten.CursorPosition()
	i.X = mx
	i.Y = my
}
