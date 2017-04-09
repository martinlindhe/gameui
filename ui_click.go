package ui

import "log"

// click handles click events for all components
func (ui *UI) handleClick() {
	if !ui.Input.NewStateForMouse(MouseButtonLeft) {
		return
	}
	mouse := Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.children {
		if c.IsHidden() {
			continue
		}
		if mouse.In(c.GetBounds()) {
			if c.Click(mouse) {
				log.Println("Click handled")
				ui.Input.ConsumeStateForMouse(MouseButtonLeft)
			}
			return
		}
	}
}

func (ui *UI) handleDrag() {
	if !ui.Input.ContinuedStateForMouse(MouseButtonLeft) {
		return
	}
	mouse := Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.children {
		if c.IsHidden() {
			continue
		}
		b := c.GetBounds()
		if c.IsDraggable() {
			if mouse.In(b) {
				log.Println("dragging...", b)
				c.Move(ui.Input.X-b.Min.X, ui.Input.Y-b.Min.Y)
				return
			}
			// log.Println("draggable but mouse is outside")
		}
	}
}
