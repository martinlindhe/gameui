package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// UI represents an instance of the UI
type UI struct {
	Width, Height int
	WindowTitle   string
	Scene         *image.RGBA
	components    []Component
	Input         Input
	keyFuncs      map[Key]func() error
}

// New creates a new UI instance
func New(width, height int) *UI {
	rect := image.Rect(0, 0, width, height)
	ui := UI{
		Width:       width,
		Height:      height,
		WindowTitle: "ui window",
		Scene:       image.NewRGBA(rect),
	}
	ui.keyFuncs = make(map[Key]func() error)
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
	ui.components = append(ui.components, o)
}

// Render returns a fresh frame of the GUI
func (ui *UI) Render(mx, my int) image.Image {
	if ui.IsClean() {
		return ui.Scene
	}
	whole := image.Rect(0, 0, ui.Width, ui.Height)
	draw.Draw(ui.Scene, whole, &image.Uniform{color.Transparent}, image.ZP, draw.Src)

	for _, c := range ui.components {
		img := c.Draw(mx, my)
		if img == nil {
			continue
		}
		r := c.GetBounds()
		c.Hover(mx >= r.Min.X && mx <= r.Max.X && my >= r.Min.Y && my <= r.Max.Y)

		draw.Draw(ui.Scene, r, img, image.ZP, draw.Over)
	}
	return ui.Scene
}

// IsClean returns true if all UI components are clean
func (ui *UI) IsClean() bool {
	for _, c := range ui.components {
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
