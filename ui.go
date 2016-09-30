package ui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hajimehoshi/ebiten"
)

// TODO rework so this component only renders a image, no loop code etc...
// do that in main app

// UI represents an instance of the UI
type UI struct {
	Width, Height  int
	scale          float64
	windowTitle    string
	elements       []Component
	screen         *ebiten.Image
	buffer         *ebiten.Image
	updateCallback func(*ebiten.Image) error
}

// New creates a new UI instance
func New(width, height int, scale float64) (*UI, error) {
	var err error
	ui := UI{
		Width:       width,
		Height:      height,
		scale:       scale,
		windowTitle: "eb-ui",
	}
	if ui.buffer, err = ebiten.NewImage(ui.Width, ui.Height, ebiten.FilterNearest); err != nil {
		return nil, err
	}
	if ui.screen, err = ebiten.NewImage(ui.Width, ui.Height, ebiten.FilterNearest); err != nil {
		return nil, err
	}
	return &ui, nil
}

// SetWindowTitle sets the title of the application window
func (ui *UI) SetWindowTitle(s string) {
	ui.windowTitle = s
}

// AddElement adds an element to the ui
func (ui *UI) AddElement(o Component) {
	ui.elements = append(ui.elements, o)
}

// Loop ...
func (ui *UI) Loop(updateCallback func(*ebiten.Image) error) error {
	if updateCallback == nil {
		log.Fatalln("api used improperly: updateCallback is nil")
	}
	ui.updateCallback = updateCallback
	if err := ebiten.Run(ui.update, ui.Width, ui.Height, ui.scale, ui.windowTitle); err != nil {
		return err
	}
	return nil
}

// callback for ebiten.Run()
func (ui *UI) update(im *ebiten.Image) error {

	err := ui.updateCallback(im)
	if err != nil {
		return err
	}

	mx, my := ebiten.CursorPosition()

	// redraw each element
	for _, el := range ui.elements {

		tile, err := el.Draw(mx, my)
		if err != nil {
			return err
		}

		ux, uy := el.GetUpperLeft()
		// log.Debugln("ebui.update: drawing tile ", i, " at ", ux, uy)

		// blend tile on image
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(ux), float64(uy))
		im.DrawImage(tile, op)
	}

	return nil
}
