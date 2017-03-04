// window containing multiple objects, click "ADD" to add more

package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	ui "github.com/martinlindhe/gameui"
)

const (
	width, height = 640, 480
	scale         = 1.
	fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
	gui       = ui.New(width, height)
	font12, _ = ui.NewFont(fontName, 12, 72, ui.White)
	font20, _ = ui.NewFont(fontName, 20, 72, ui.White)
	fps       = ui.NewText(font20)
	count     = uint64(0)
)

// obj implements IconGroupObject interface
type obj struct {
	name    string
	id      uint64
	icon    *image.RGBA
	onClick func(*obj)
}

func (o obj) Name() string {
	return o.name
}
func (o obj) ID() uint64 {
	return o.id
}
func (o obj) Icon() image.Image {
	return o.icon
}
func (o obj) Click() {
	o.onClick(&o)
}

func main() {
	wnd := makeWindow()
	gui.AddComponent(wnd)
	gui.AddComponent(fps)

	if err := ebiten.Run(update, width, height, scale, "Dialog (UI Demo)"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if err := gui.Update(); err != nil {
		return err
	}

	fps.SetText(fmt.Sprintf("%.1f", ebiten.CurrentFPS()))
	frame, err := ebiten.NewImageFromImage(gui.Render(0, 0), ebiten.FilterNearest)
	if err != nil {
		return err
	}
	screen.DrawImage(frame, &ebiten.DrawImageOptions{})
	return nil
}

func makeWindow() *ui.Window {
	iconW, iconH := 3, 3
	wnd := ui.NewWindow(width-30, height-30, "icon group")
	wnd.Position = ui.Point{X: 15, Y: 25}

	grp := ui.NewIconGroup(5, 5, iconW, iconH)
	grp.Position = ui.Point{X: 0, Y: 0} // XXX should become relative to y=0 = after title bar
	wnd.AddChild(grp)

	btnAdd := ui.NewButton(60, 22).
		SetText(font12, "ADD")
	btnAdd.Position = ui.Point{
		X: 0,
		Y: wnd.Dimension.Height - btnAdd.Dimension.Height - wnd.TitlebarHeight(),
	}
	btnAdd.OnClick = func() {
		name := "icon " + fmt.Sprintf("%d", count)
		fmt.Println("adding obj", name)
		im1 := image.NewRGBA(image.Rect(0, 0, iconW, iconH))
		im1.Set(0, 0, ui.White)
		im1.Set(2, 0, ui.White)
		im1.Set(1, 2, ui.White)
		o := obj{name: name, id: count, icon: im1, onClick: func(o *obj) {
			fmt.Println("CLICKED", o.name, "idx", o.id, "so we remove it")
			grp.RemoveObjectByID(o.id)
		}}
		grp.AddObject(o)
		count++
	}
	wnd.AddChild(btnAdd)

	btnBye := ui.NewButton(60, 22).
		SetText(font12, "BYE")
	btnBye.Position = ui.Point{
		X: wnd.Dimension.Width - btnBye.Dimension.Width,
		Y: wnd.Dimension.Height - btnBye.Dimension.Height - wnd.TitlebarHeight(),
	}
	btnBye.OnClick = func() {
		fmt.Println("exiting")
		os.Exit(0)
	}
	wnd.AddChild(btnBye)

	return wnd
}
