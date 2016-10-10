// XXX TODO window containing multiple objects of loose interface with just "Name() string" identifier

// XXX TODO later: resizable window
// XXX TODO later: minimize _ icon on window
// XXX TODO window title

package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/martinlindhe/farm/game"
	"github.com/martinlindhe/farm/ui"
)

const (
	width, height = 320, 200
)

var (
	gui = ui.New(width, height)
	fps = ui.NewText(30, color.White)
)

func main() {

	wnd := makeWindow()
	gui.AddComponent(wnd)

	gui.AddComponent(fps)

	if err := ebiten.Run(update, width, height, 1, "Dialog (UI Demo)"); err != nil {
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
	wnd := ui.NewWindow(200, 100).
		SetTitle("icon group")
	wnd.Position = ui.Point{X: 10, Y: 10}

	// XXX each icon of interface type to get Name() and Icon()
	exit := ui.NewIconGroup(16, 16, game.ItemTileW, game.ItemTileH)
	exit.Position = ui.Point{X: width / 3, Y: height / 3}
	wnd.AddChild(exit)

	btnYes := ui.NewButton(60, 20).
		SetText("YES")
	btnYes.Position = ui.Point{X: 0, Y: height / 2}
	btnYes.OnClick = func() {
		fmt.Println("clicked", btnYes.Text.GetText())
		os.Exit(0)
	}
	wnd.AddChild(btnYes)

	btnNo := ui.NewButton(60, 20).
		SetText("NO")
	btnNo.Position = ui.Point{X: width / 2, Y: height / 2}
	btnNo.OnClick = func() {
		fmt.Println("clicked", btnNo.Text.GetText())
	}
	wnd.AddChild(btnNo)
	return wnd
}
