package ui

// click handles click events for all components
func (ui *UI) click() {
	if !ui.Input.StateForMouse(MouseButtonLeft) {
		return
	}

	mouse := Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.components {
		if mouse.In(c.GetBounds()) {
			c.Click(mouse)
			return
		}
	}
}
