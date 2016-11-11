package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Window is a window (UI component)
type Window struct {
	component
	titlebarHeight  int
	childLeftPad    int
	titleColor      color.Color
	backgroundColor color.Color
	borderColor     color.Color
	title           *Text
	close           *Button
}

var (
	windowBgColor     = color.RGBA{0x50, 0x50, 0x50, 192} // gray, 25% transparent
	windowBorderColor = color.RGBA{0x40, 0x40, 0x40, 192} // white, 25% transparent
	windowTitleColor  = color.RGBA{0x50, 0x50, 0x50, 192} //gray
)

// NewWindow creates a new window with height of `height` + height of title bar
func NewWindow(width, height int, titleText string) *Window {
	wnd := Window{}
	wnd.Dimension.Width = width
	wnd.Dimension.Height = height

	wnd.backgroundColor = windowBgColor
	wnd.borderColor = windowBorderColor
	wnd.titleColor = windowTitleColor

	title := NewText(10, White)
	title.Position = Point{X: 1, Y: 0}
	wnd.title = title
	wnd.title.SetText(titleText)
	wnd.addChild(title)

	if titleText != "" {
		wnd.titlebarHeight = wnd.title.GetHeight() + 1
		wnd.childLeftPad = 1
	}

	close := NewButton(wnd.titlebarHeight, wnd.titlebarHeight)
	close.OnClick = func() {
		wnd.Hide()
	}
	wnd.close = close
	wnd.addChild(close)

	return &wnd
}

// AddChild adds a child component to the window, adjusting position
func (wnd *Window) AddChild(c Component) {
	if val, ok := c.(Positioner); ok {
		pos := val.GetPosition()
		pos.X += wnd.childLeftPad
		pos.Y += wnd.titlebarHeight
		val.SetPosition(pos)
	}
	wnd.addChild(c)
}

func (wnd *Window) addChild(c Component) {
	wnd.isClean = false
	wnd.children = append(wnd.children, c)
}

// HideCloseButton ...
func (wnd *Window) HideCloseButton() {
	wnd.close.Hide()
}

// HideTitle ...
func (wnd *Window) HideTitle() {
	wnd.title.Hide()
}

// SetTitle ...
func (wnd *Window) SetTitle(s string) *Window {
	wnd.title.SetText(s)
	return wnd
}

// SetTitleColor ...
func (wnd *Window) SetTitleColor(col color.Color) {
	wnd.titleColor = col
}

// SetBackgroundColor ...
func (wnd *Window) SetBackgroundColor(col color.Color) {
	wnd.backgroundColor = col
}

// SetBorderColor ...
func (wnd *Window) SetBorderColor(col color.Color) {
	wnd.borderColor = col
}

// Draw redraws internal buffer
func (wnd *Window) Draw(mx, my int) *image.RGBA {
	if wnd.isHidden {
		return nil
	}
	if wnd.isClean {
		if !wnd.isChildrenClean() {
			wnd.isClean = false
		} else {
			return wnd.Image
		}
	}
	wnd.initImage()

	wnd.close.Position = Point{X: wnd.Dimension.Width - wnd.close.Dimension.Width, Y: 0}

	// draw background color
	rect := image.Rect(0, 0, wnd.Dimension.Width, wnd.Dimension.Height)
	draw.Draw(wnd.Image, rect, &image.Uniform{wnd.backgroundColor}, image.ZP, draw.Over)

	// draw titlebar rect
	if !wnd.title.isHidden {
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
