package ui

import "log"

// click handles click events for all components
func (ui *UI) handleClick() {
	if !ui.Input.StateForMouse(MouseButtonLeft) {
		return
	}
	mouse := Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.children {
		if c.IsHidden() {
			continue
		}
		if mouse.In(c.GetBounds()) {
			if c.Click(mouse) {
				ui.Input.ConsumeStateForMouse(MouseButtonLeft)
			}
			return
		}
	}
}

func (ui *UI) handleDrag() {
	if !ui.Input.StateForMouse(MouseButtonLeft) {
		return
	}
	mouse := Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.children {
		if c.IsHidden() {
			continue
		}
		if c.IsDraggable() && mouse.In(c.GetBounds()) {
			log.Println("dragging...")
			c.Move(ui.Input.X, ui.Input.Y)

			//ui.Input.ConsumeStateForMouse(MouseButtonLeft)
			return
		}
	}
}
