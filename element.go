package ui

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hajimehoshi/ebiten"
)

// Component represents any type of UI component
type Component interface {
	Draw(mx, my int) (*ebiten.Image, error)
	GetUpperLeft() (int, int)
}

type element struct {
	IsHovering    bool
	Width, Height int
	X, Y          int
	Image         *ebiten.Image
}

func (el element) GetUpperLeft() (int, int) {
	return el.X, el.Y
}

// updateHover toggles IsHovering if cursor is over element
func (el element) updateHover(mx, my int) {

	el.IsHovering = false
	if mx >= el.X && mx <= el.X+el.Width &&
		my >= el.Y && my <= el.Y+el.Height {
		el.IsHovering = true
	}

	log.Debugln("element.updateHover:", el.X, el.Y, el.IsHovering)
}
