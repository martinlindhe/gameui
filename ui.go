package ui

import (
	"image"
	"image/draw"
)

// UI represents an instance of the UI
type UI struct {
	Width, Height int
	scale         float64
	WindowTitle   string
	Scene         *image.RGBA // XXX unused now. TODO: should contain last rendered scene
	components    []Component
}

// New creates a new UI instance
func New(width, height int) *UI {
	rect := image.Rect(0, 0, width, height)
	return &UI{
		Width:       width,
		Height:      height,
		WindowTitle: "ui window",
		Scene:       image.NewRGBA(rect),
	}
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
func (ui *UI) Render() *image.RGBA {
	// XXX if all components are clean, reuse last drawn frame

	dst := image.NewRGBA(image.Rect(0, 0, ui.Width, ui.Height))

	for _, c := range ui.components {
		img := c.Draw()
		x, y, w, h := c.GetBounds()
		dr := image.Rect(x, y, x+w, y+h)
		draw.Draw(dst, dr, img, image.ZP, draw.Over)
	}
	return dst
}
