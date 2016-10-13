package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Window ...
type Window struct {
	component
	titleColor      color.Color
	backgroundColor color.Color
	borderColor     color.Color
	title           *Text
	close           *Button
}

var (
	windowBgColor     = color.RGBA{0x50, 0x50, 0x50, 192} // gray, 75% transparent
	windowBorderColor = White
	windowTitleColor  = color.RGBA{0x50, 0x50, 0x50, 255} //gray
)

// NewWindow ...
func NewWindow(width, height int) *Window {
	wnd := Window{}
	wnd.Dimension.Width = width
	wnd.Dimension.Height = height

	wnd.backgroundColor = windowBgColor
	wnd.borderColor = windowBorderColor
	wnd.titleColor = windowTitleColor

	title := NewText(12, White)
	title.Position = Point{X: 0, Y: 0}
	wnd.title = title
	wnd.AddChild(title)

	close := NewButton(10, 10)
	close.Position = Point{X: wnd.Dimension.Width - close.GetBounds().Max.X, Y: 0}
	close.OnClick = func() {
		wnd.Hide()
	}
	wnd.close = close
	wnd.AddChild(close)

	return &wnd
}

// AddChild ...
func (wnd *Window) AddChild(c Component) {
	wnd.children = append(wnd.children, c)
	wnd.isClean = false
}

// HideCloseButton ...
func (wnd *Window) HideCloseButton(b bool) {
	wnd.close.Hidden = b
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
	if wnd.Hidden {
		return nil
	}
	if wnd.isClean {
		if !wnd.isChildrenClean() {
			wnd.isClean = false
		} else {
			return wnd.Image
		}
	}

	rect := image.Rect(0, 0, wnd.Dimension.Width, wnd.Dimension.Height)
	if wnd.Image == nil {
		wnd.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(wnd.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}

	// draw background color
	draw.Draw(wnd.Image, rect, &image.Uniform{wnd.backgroundColor}, image.ZP, draw.Over)

	titlebarH := wnd.title.Dimension.Height + 1

	// draw titlebar rect
	titleRect := image.Rect(0, 0, wnd.Dimension.Width, titlebarH)
	draw.Draw(wnd.Image, titleRect, &image.Uniform{wnd.titleColor}, image.ZP, draw.Over)

	// draw outline
	DrawRect(wnd.Image, &rect, wnd.borderColor)

	wnd.drawChildren(mx, my)

	wnd.isClean = true
	return wnd.Image
}

// Click pass click to window child components
func (wnd *Window) Click(mouse Point) {
	childPoint := Point{X: mouse.X - wnd.Position.X, Y: mouse.Y - wnd.Position.Y}
	for _, c := range wnd.children {
		if childPoint.In(c.GetBounds()) {
			c.Click(childPoint)
			return
		}
	}
}
