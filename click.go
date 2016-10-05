package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Click handles click events for all components
func (ui *UI) Click() {
	ui.Input.UpdateMouse()
	if !ui.Input.StateForMouse(ebiten.MouseButtonLeft) {
		return
	}

	mouse := image.Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.components {
		if mouse.In(c.GetRect()) {
			c.Click()
			return
		}
	}
}
