package ui

import (
	"image"
	"image/color"
	"image/draw"
	"log"
)

// Window is a window (UI component)
type Window struct {
	component
	titlebarHeight int
	childLeftPad   int
	titleColor     color.Color
	borderColor    color.Color
	title          *Text
	close          *Button
}

var (
	windowBgColor     = color.RGBA{0x50, 0x50, 0x50, 192} // grey, 25% transparent
	windowBorderColor = color.RGBA{0x40, 0x40, 0x40, 192} // white, 25% transparent
	windowTitleColor  = color.RGBA{0x50, 0x50, 0x50, 192} // grey
)

// NewWindow creates a new window with height of `height` + height of title bar
func NewWindow(width, height int, titleText string) *Window {
	wnd := Window{}
	wnd.Dimension.Width = width
	wnd.Dimension.Height = height

	wnd.backgroundColor = windowBgColor
	wnd.borderColor = windowBorderColor
	wnd.titleColor = windowTitleColor

	if titleText != "" {
		wnd.SetTitle(titleText)
	}

	close := NewButton(wnd.titlebarHeight, wnd.titlebarHeight)
	close.OnClick = func() {
		wnd.Hide()
	}
	wnd.close = close
	wnd.addChild(close)

	return &wnd
}

// AddChild adds a child to the window
func (wnd *Window) AddChild(c Component) {
	wnd.addChild(c)
}

// XXX:
/*
func (wnd *Window) AddChild(c Component) {
	if val, ok := c.(Positioner); ok {
		pos := val.GetPosition()
		pos.X += wnd.childLeftPad
		pos.Y += wnd.titlebarHeight
		val.SetPosition(pos)
	}
	wnd.addChild(c)
}
*/

// HideCloseButton ...
func (wnd *Window) HideCloseButton() {
	wnd.close.Hide()
}

// HideTitle ...
func (wnd *Window) HideTitle() {
	if wnd.title != nil {
		wnd.title.Hide()
	}
}

// SetTitle ...
func (wnd *Window) SetTitle(s string) *Window {
	if s == "" {
		return wnd
	}
	fnt, err := NewFont(defaultFontName, 10, 72, White)
	if err != nil {
		log.Println("ERROR:", err)
		return wnd
	}
	title := NewText(fnt)
	title.Position = Point{X: 1, Y: 0}
	wnd.title = title
	wnd.title.SetText(s)
	wnd.addChild(title)

	wnd.titlebarHeight = wnd.title.GetHeight() + 1
	wnd.childLeftPad = 1
	return wnd
}

// SetTitleColor ...
func (wnd *Window) SetTitleColor(col color.Color) {
	wnd.titleColor = col
}

// SetBorderColor ...
func (wnd *Window) SetBorderColor(col color.Color) {
	wnd.borderColor = col
}

// Draw redraws internal buffer
func (wnd *Window) Draw(mx, my int) *image.RGBA {
	if wnd.isHidden {
		wnd.isClean = true
		return nil
	}
	if wnd.isClean {
		if wnd.isChildrenClean() {
			return wnd.Image
		}
	}
	wnd.initImage()

	wnd.close.Position = Point{X: wnd.Dimension.Width - wnd.close.Dimension.Width, Y: 0}

	// draw background color
	rect := image.Rect(0, 0, wnd.Dimension.Width, wnd.Dimension.Height)
	draw.Draw(wnd.Image, rect, &image.Uniform{wnd.backgroundColor}, image.ZP, draw.Over)

	// draw titlebar rect
	if wnd.title != nil && !wnd.title.isHidden {
		titleRect := image.Rect(0, 0, wnd.Dimension.Width, wnd.titlebarHeight)
		draw.Draw(wnd.Image, titleRect, &image.Uniform{wnd.titleColor}, image.ZP, draw.Over)
	}

	// draw outline
	outlineRect := image.Rect(0, 0, wnd.Dimension.Width-1, wnd.Dimension.Height-1)
	DrawRect(wnd.Image, outlineRect, wnd.borderColor)

	wnd.drawChildren(mx, my)

	wnd.isClean = true
	return wnd.Image
}

// Click pass click to window child components
func (wnd *Window) Click(mouse Point) bool {
	childPoint := Point{X: mouse.X - wnd.Position.X, Y: mouse.Y - wnd.Position.Y}
	for _, c := range wnd.children {
		if childPoint.In(c.GetBounds()) {
			c.Click(childPoint)
			return true
		}
	}
	return false
}

// TitlebarHeight ...
func (wnd *Window) TitlebarHeight() int {
	return wnd.titlebarHeight
}
