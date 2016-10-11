package ui

import (
	"image"
	"image/color"
	"image/draw"
)

// Window ...
type Window struct {
	component
	title           string
	titleColor      color.Color
	backgroundColor color.Color
}

var (
	windowTitleColor = color.RGBA{0x50, 0x50, 0x50, 255} //gray
	windowBgColor    = color.RGBA{0x50, 0x50, 0x50, 192} // gray, 75% transparent
)

// NewWindow ...
func NewWindow(width, height int) *Window {
	wnd := Window{title: "new window"}
	wnd.Dimension.Width = width
	wnd.Dimension.Height = height
	wnd.titleColor = windowTitleColor
	wnd.backgroundColor = windowBgColor
	return &wnd
}

// AddChild ...
func (wnd *Window) AddChild(c Component) {
	wnd.children = append(wnd.children, c)
}

// SetTitle ...
func (wnd *Window) SetTitle(s string) *Window {
	wnd.title = s
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

	/* XXX add a child font object for this:
	title := "BLUEPRINTS"
	if err := common.ArcadeFont.DrawTextOnImage(wnd.Image, title, x0+1, y0+1); err != nil {
		panic(err)
	}
	*/

	// draw outline
	DrawRect(wnd.Image, &rect, color.White)

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
