package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// UI represents an instance of the UI
type UI struct {
	Width, Height int
	scale         float64
	WindowTitle   string
	Scene         *image.RGBA
	components    []Component
	Input         Input
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
func (ui *UI) Render(mx, my int) *image.RGBA {

	// do any component need to be redrawn?
	dirty := false
	for _, c := range ui.components {
		if !c.IsClean() {
			dirty = true
			break
		}
	}
	if !dirty {
		return ui.Scene
	}

	// clear scene
	whole := image.Rect(0, 0, ui.Width, ui.Height)
	draw.Draw(ui.Scene, whole, &image.Uniform{color.Transparent}, image.ZP, draw.Src)

	for _, c := range ui.components {
		img := c.Draw(mx, my)
		x, y, w, h := c.GetBounds()
		x1 := x + w
		y1 := y + h
		c.Hover(mx >= x && mx <= x1 && my >= y && my <= y1)

		dr := image.Rect(x, y, x1, y1)
		draw.Draw(ui.Scene, dr, img, image.ZP, draw.Over)
	}
	return ui.Scene
}
