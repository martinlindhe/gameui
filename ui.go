package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// UI represents an instance of the UI
type UI struct {
	component
	WindowTitle  string
	Input        Input
	keyFuncs     map[Key]func() error
	prevX, prevY int
}

// New creates a new UI instance
func New(width, height int) *UI {
	rect := image.Rect(0, 0, width, height)
	ui := UI{
		WindowTitle: "ui window",
		keyFuncs:    make(map[Key]func() error),
	}
	ui.Dimension = Dimension{Width: width, Height: height}
	ui.Image = image.NewRGBA(rect)
	return &ui
}

// AddKeyFunc registers a function to run on key press
func (ui *UI) AddKeyFunc(key Key, fnc func() error) {
	ui.keyFuncs[key] = fnc
}

// Update is called on every frame from the ebiten.Run update callback
func (ui *UI) Update() error {
	ui.Input.updateMouse()
	ui.handleClick()
	if err := ui.handleKeypress(); err != nil {
		return err
	}
	return nil
}

// SetWindowTitle sets the title of the application window
func (ui *UI) SetWindowTitle(s string) {
	ui.WindowTitle = s
}

// AddComponent adds a component to the ui
func (ui *UI) AddComponent(o Component) {
	ui.addChild(o)
}

// Render returns a fresh frame of the GUI. mx, my is absolute mouse position
func (ui *UI) Render() image.Image {
	mx := ui.Input.X
	my := ui.Input.Y
	if ui.isChildrenClean() {
		if mx == ui.prevX && my == ui.prevY {
			return ui.Image
		}
		if (mx < 0 || mx > ui.Dimension.Width) || (my < 0 || my > ui.Dimension.Height) {
			// cursor outside window will not change hover state
			return ui.Image
		}
	}
	ui.prevX = mx
	ui.prevY = my
	whole := image.Rect(0, 0, ui.Dimension.Width, ui.Dimension.Height)
	draw.Draw(ui.Image, whole, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	ui.drawChildren(mx, my)
	ui.drawTooltips(mx, my)
	return ui.Image
}

// mx, my is absolute mouse position
func (ui *UI) drawTooltips(mx, my int) {
	for _, child := range ui.children {
		if grp, ok := child.(*Group); ok {
			for _, gchild := range grp.children {
				ui.drawTooltip(gchild, mx, my, mx-grp.Position.X, my-grp.Position.Y)
			}
		}
		ui.drawTooltip(child, mx, my, mx, my)
	}
}

func (ui *UI) drawTooltip(child Component, mx, my, relx, rely int) {
	r := child.GetBounds()
	child.Hover(relx >= r.Min.X && relx <= r.Max.X && rely >= r.Min.Y && rely <= r.Max.Y)

	tooltip := child.Tooltip()
	if child.IsMouseOver() && tooltip != nil {
		tooltip.Move(mx, my)
		tr := tooltip.GetBounds()
		draw.Draw(ui.Image, tr, tooltip.Draw(relx, rely), image.ZP, draw.Over)
	}
}

// IsClean returns true if all UI components are clean
func (ui *UI) IsClean() bool {
	for _, c := range ui.children {
		if !c.IsClean() {
			return false
		}
	}
	return true
}

// handleKeypress runs corresponding function when a key is pressed
func (ui *UI) handleKeypress() error {
	ui.Input.updateKeyboard()
	for key, fnc := range ui.keyFuncs {
		if ui.Input.StateForKey(key) {
			if err := fnc(); err != nil {
				return err
			}
		}
	}
	return nil
}
