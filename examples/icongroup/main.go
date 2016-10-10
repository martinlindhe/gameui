// window containing multiple objects, click "ADD" to add more

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/martinlindhe/farm/ui"
)

const (
	width, height = 320, 200
	scale         = 2.
)

var (
	gui = ui.New(width, height)
	fps = ui.NewText(30, color.White)
)

// obj implements IconGroupObject interface
type obj struct {
	name string
	icon *image.RGBA
}

func (o obj) Name() string {
	return o.name
}
func (o obj) Icon() *image.RGBA {
	return o.icon
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
	if err := screen.DrawImage(frame, &ebiten.DrawImageOptions{}); err != nil {
		return err
	}
	return nil
}

func makeWindow() *ui.Window {
	wnd := ui.NewWindow(width-30, height-30).
		SetTitle("icon group")
	wnd.Position = ui.Point{X: 15, Y: 25}

	grp := ui.NewIconGroup(5, 5, 3, 3)
	grp.Position = ui.Point{X: 0, Y: 0}
	wnd.AddChild(grp)

	btnBye := ui.NewButton(60, 20).
		SetText("BYE")
	btnBye.Position = ui.Point{X: 0, Y: wnd.Height - 20}
	btnBye.OnClick = func() {
		fmt.Println("exiting")
		os.Exit(0)
	}
	wnd.AddChild(btnBye)

	btnAdd := ui.NewButton(60, 20).
		SetText("ADD")
	btnAdd.Position = ui.Point{X: wnd.Width / 2, Y: wnd.Height - 20}
	btnAdd.OnClick = func() {
		fmt.Println("adding obj")
		im1 := image.NewRGBA(image.Rect(0, 0, 3, 3))
		im1.Set(0, 0, color.White)
		im1.Set(2, 0, color.White)
		im1.Set(1, 2, color.White)
		grp.AddObject(obj{name: "icon", icon: im1})
	}
	wnd.AddChild(btnAdd)
	return wnd
}
