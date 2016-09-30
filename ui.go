package ui

import (
	"image"
)

// UI represents an instance of the UI
type UI struct {
	Width, Height int
	scale         float64
	WindowTitle   string
	//elements      []Component
	screen image.Image
	buffer image.Image
}

// New creates a new UI instance
func New(width, height int) *UI {
	rect := image.Rect(0, 0, width, height)
	return &UI{
		Width:       width,
		Height:      height,
		WindowTitle: "ui window",
		buffer:      image.NewRGBA(rect),
		screen:      image.NewRGBA(rect),
	}
}

// SetWindowTitle sets the title of the application window
func (ui *UI) SetWindowTitle(s string) {
	ui.WindowTitle = s
}

/*
// AddElement adds an element to the ui
func (ui *UI) AddElement(o Component) {
	ui.elements = append(ui.elements, o)
}
*/
