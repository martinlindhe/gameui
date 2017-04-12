package ui

// click handles click events for all components, returns true if click was handled
func (ui *UI) handleClick() bool {
	if !ui.Input.NewStateForMouse(MouseButtonLeft) {
		return false
	}
	mouse := Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.children {
		if c.IsHidden() {
			continue
		}
		if mouse.In(c.GetBounds()) {
			if c.Click(mouse) {
				//log.Println("Click handled")
				ui.Input.ConsumeStateForMouse(MouseButtonLeft)
				return true
			}

			b := c.GetBounds()
			if c.IsDraggable() && !c.IsDragged() && mouse.In(b) && !ui.dragStarting {
				//log.Println("drag = true")
				ui.dragStarting = true
				ui.dragX = ui.Input.X
				ui.dragY = ui.Input.Y
				return false

			}
		}
	}
	return false
}

func (ui *UI) handleDrag() {
	canDrag := ui.Input.ContinuedStateForMouse(MouseButtonLeft)
	mouse := Point{X: ui.Input.X, Y: ui.Input.Y}
	for _, c := range ui.children {
		if c.IsHidden() {
			continue
		}

		b := c.GetBounds()

		if canDrag && c.IsDraggable() && !c.IsDragged() && mouse.In(b) && ui.dragStarting && (ui.dragX != ui.Input.X || ui.dragY != ui.Input.Y) {
			//log.Println("Drag started drag=", ui.dragX, ui.dragY, "now=", ui.Input.X, ui.Input.Y)
			c.Drag(true)
		}

		if canDrag && c.IsDragged() {
			//log.Println("dragging", ui.Input.X, b.Min.X, ui.Input.X-b.Min.X)
			//c.Move(b.Min.X+(ui.Input.X-b.Min.X), b.Min.Y+(ui.Input.Y-b.Min.Y))

			relX, relY := ui.Input.X-b.Min.X, ui.Input.Y-b.Min.Y
			dstX, dstY := b.Min.X+relX, b.Min.Y+relY // XXX in effect just move to mouse point, which is wrong
			//dstX, dstY := relX, relY
			//dstX, dstY := ui.Input.X-relX, ui.Input.Y-relY // feels right but doesnt move at all

			//log.Println("dragging to", dstX, dstY, "obj is at", b.Min.X, b.Min.Y, "mouse at", ui.Input.X, ui.Input.Y, "relative to child", relX, relY)
			c.Move(dstX, dstY)
		}

		if !canDrag {
			//log.Println("drag released")
			c.Drag(false)
			ui.dragStarting = false
		}
	}
}
