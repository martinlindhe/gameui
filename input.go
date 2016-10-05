package ui

// XXX can this be done without ebiten?

import "github.com/hajimehoshi/ebiten"

// parts from ebiten/examples/blocks/blocks/input.go

// Input ...
type Input struct {
	keyStates   [256]int
	mouseStates [3]int
}

// StateForKey ...
func (i *Input) StateForKey(key ebiten.Key) bool {
	return i.keyStates[key] == 1
}

// StateForMouse ...
func (i *Input) StateForMouse(mouse ebiten.MouseButton) bool {
	return i.mouseStates[mouse] == 1
}

// ClearMouse clear the mouse state
func (i *Input) ClearMouse() {
	for mouse := range i.mouseStates {
		i.mouseStates[mouse] = 0
	}
}

// UpdateKeyboard ...
func (i *Input) UpdateKeyboard() {
	for key := range i.keyStates {
		if !ebiten.IsKeyPressed(ebiten.Key(key)) {
			i.keyStates[key] = 0
			continue
		}
		i.keyStates[key]++
	}
}

// UpdateMouse ...
func (i *Input) UpdateMouse() {
	for mouse := range i.mouseStates {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButton(mouse)) {
			i.mouseStates[mouse] = 0
			continue
		}
		i.mouseStates[mouse]++
	}
}
