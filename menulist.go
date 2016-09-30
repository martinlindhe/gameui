package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var (
	pad           = 2.0
	menuListColor = color.RGBA{209, 154, 86, 255}
)

// MenuList represents a ordered list of MenuItem:s
type MenuList struct {
	element
	Items []MenuItem
}

// NewMenuList ...
func NewMenuList(ebui *UI) *MenuList {

	o := new(MenuList)
	o.Width = ebui.Width / 2
	return o
}

// AddItem ...
func (el *MenuList) AddItem(item *MenuItem) {

	el.Items = append(el.Items, *item)
}

// Draw renders the MenuList with all children
func (el MenuList) Draw(mx, my int) (*ebiten.Image, error) {

	el.Image.Fill(menuListColor)

	for i, item := range el.Items {

		childEl, err := item.Draw(mx, my)
		if err != nil {
			return nil, err
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pad, pad+float64(i*(menuItemHeight+int(pad))))
		el.Image.DrawImage(childEl, op)
	}

	return el.Image, nil
}
