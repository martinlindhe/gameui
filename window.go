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
	windowBorderColor = color.White
	windowTitleColor  = color.RGBA{0x50, 0x50, 0x50, 255} //gray
)

// NewWindow ...
func NewWindow(width, height int) *Window {
	wnd := Window{}
	wnd.title = NewText(12, color.White)
	wnd.close = NewButton(10, 10)
	wnd.Dimension.Width = width
	wnd.Dimension.Height = height
	wnd.backgroundColor = windowBgColor
	wnd.borderColor = windowBorderColor
	wnd.titleColor = windowTitleColor
	return &wnd
}

// AddChild ...
func (wnd *Window) AddChild(c Component) {
	wnd.children = append(wnd.children, c)
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

	textH := 10 // XXX
	titlebarH := textH + 1

	// draw titlebar rect
	titleRect := image.Rect(0, 0, wnd.Dimension.Width, titlebarH)
	draw.Draw(wnd.Image, titleRect, &image.Uniform{wnd.titleColor}, image.ZP, draw.Over)

	// draw titlebar text
	title := wnd.title.Draw(mx, my)
	draw.Draw(wnd.Image, title.Bounds(), title, image.ZP, draw.Over)

	closeBnt := wnd.close.Draw(mx, my)
	draw.Draw(wnd.Image, closeBnt.Bounds(), closeBnt, image.ZP, draw.Over)

	//fmt.Println("first", ax0, ay0, "second", ax1, ay1)
	/* XXX handle click X
	if isPointInsideRect(mouse, &closeRect) &&
		r.world.Input.StateForMouse(ebiten.MouseButtonLeft) {
		fmt.Println("XXX clicked close X")
		r.ShowBuildMenu = false
		return
	}
	*/

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
