package ui

import "image"

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
