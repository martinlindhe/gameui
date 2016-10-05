package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Click handles click events for all components
func (ui *UI) Click(mouse *Point) {
	ui.Input.UpdateMouse()
	if !ui.Input.StateForMouse(ebiten.MouseButtonLeft) {
		return
	}

	for _, c := range ui.components {
		if image.Point(*mouse).In(c.GetRect()) {
			c.Click()
			return
		}
	}
}
