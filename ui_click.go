package ui

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
