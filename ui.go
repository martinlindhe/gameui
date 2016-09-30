package ui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hajimehoshi/ebiten"
)

// TODO rework so this component only renders a image, no loop code etc... do that in main app

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

	ebui := UI{
		Width:       width,
		Height:      height,
		scale:       scale,
		windowTitle: "eb-ui",
	}

	if ebui.buffer, err = ebiten.NewImage(ebui.Width, ebui.Height, ebiten.FilterNearest); err != nil {
		return nil, err
	}

	if ebui.screen, err = ebiten.NewImage(ebui.Width, ebui.Height, ebiten.FilterNearest); err != nil {
		return nil, err
	}

	return &ebui, nil
}

// SetWindowTitle sets the title of the application window
func (ebui *UI) SetWindowTitle(s string) {
	ebui.windowTitle = s
}

// AddElement adds an element to the ui
func (ebui *UI) AddElement(o Component) error {

	switch el := o.(type) {
	case *MenuList:
		// calculate absolute position for all childs of `el`
		for i, item := range el.Items {
			height := (i * (menuItemHeight + int(pad)))
			item.X = el.X + item.X + int(pad)
			item.Y = el.Y + height + item.Y + int(pad)
			item.Width = el.Width - int(pad) - int(pad)
			el.Items[i] = item
		}

		el.Height = (len(el.Items) * (menuItemHeight + int(pad))) + int(pad)

		var err error
		el.Image, err = ebiten.NewImage(el.Width, el.Height, ebiten.FilterNearest)
		if err != nil {
			panic(err)
		}
	default:
		log.Println("unhandled type ", el)
	}

	ebui.elements = append(ebui.elements, o)
	return nil
}

// Loop ...
func (ebui *UI) Loop(updateCallback func(*ebiten.Image) error) error {
	if updateCallback == nil {
		log.Fatalln("api used improperly: updateCallback is nil")
	}
	ebui.updateCallback = updateCallback
	if err := ebiten.Run(ebui.update, ebui.Width, ebui.Height, ebui.scale, ebui.windowTitle); err != nil {
		return err
	}
	return nil
}

// callback for ebiten.Run()
func (ebui *UI) update(im *ebiten.Image) error {

	err := ebui.updateCallback(im)
	if err != nil {
		return err
	}

	mx, my := ebiten.CursorPosition()

	// redraw each element
	for _, el := range ebui.elements {

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
