package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	menuItemHeight    = 18
	menuItemColor     = color.RGBA{229, 192, 98, 255}
	menuItemHighlight = color.RGBA{249, 222, 198, 255}
	menuItemDisable   = color.RGBA{128, 128, 128, 255}
)

// MenuItem represents a menu item
type MenuItem struct {
	element
	Title    string
	Disabled bool
	Action   func()
	font     *Font
}

// NewMenuItem initializes the component
func NewMenuItem(ebui *UI) *MenuItem {

	o := new(MenuItem)
	var err error
	o.Width = (ebui.Width / 2) - int(pad) - int(pad)
	o.Height = menuItemHeight

	o.font, err = NewFont("_resources/font/topaz-8.ttf", 72, 12)
	if err != nil {
		panic(err)
	}

	o.Image, err = ebiten.NewImage(o.Width, o.Height, ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	return o
}

// Draw repaints the component
func (el MenuItem) Draw(mx, my int) (*ebiten.Image, error) {

	el.IsHovering = false
	if !el.Disabled &&
		mx >= el.X && mx <= el.X+el.Width &&
		my >= el.Y && my <= el.Y+el.Height {
		el.IsHovering = true
	}

	var err error

	if el.Disabled {
		el.Image.Fill(menuItemDisable)
	} else if el.IsHovering {
		el.Image.Fill(menuItemHighlight)
	} else {
		el.Image.Fill(menuItemColor)
	}

	txt, err := el.font.Print(el.Title)
	if err != nil {
		return nil, err
	}

	op := &ebiten.DrawImageOptions{}
	el.Image.DrawImage(txt, op)

	if el.IsHovering && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if el.Action == nil {
			panic("api misuse: el.Action is nil")
		}
		el.Action()
	}

	return el.Image, nil
}
