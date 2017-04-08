// a dialog window, with a yes & no button

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	ui "github.com/martinlindhe/gameui"
)

const (
	width, height = 320, 200
	scale         = 1.
	fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
	gui       = ui.New(width, height)
	font12, _ = ui.NewFont(fontName, 12, 72, ui.White)
	font20, _ = ui.NewFont(fontName, 20, 72, ui.White)
	fps       = ui.NewText(font12)
)

func main() {
	exit := ui.NewText(font20).SetText("exit?")
	exit.Position = ui.Point{X: width/2 - exit.GetWidth()/2, Y: height / 3}
	gui.AddComponent(exit)

	btnYes := ui.NewButton(60, 20).SetText(font12, "YES")
	btnYes.Position = ui.Point{X: width/4 - btnYes.Dimension.Width/2, Y: height / 2}
	btnYes.OnClick = func() {
		fmt.Println("clicked", btnYes.Text.GetText())
		os.Exit(0)
	}
	gui.AddComponent(btnYes)

	btnNo := ui.NewButton(60, 20).SetText(font12, "NO")
	btnNo.Position = ui.Point{X: (width/4)*3 - btnYes.Dimension.Width/2, Y: height / 2}
	btnNo.OnClick = func() {
		fmt.Println("clicked", btnNo.Text.GetText())
	}
	gui.AddComponent(btnNo)

	gui.AddComponent(fps)

	gui.AddKeyFunc(ui.KeyQ, func() error {
		fmt.Println("q - QUITTING")
		return ui.GracefulExitError{}
	})

	if err := ebiten.Run(update, width, height, scale, "Dialog (UI Demo)"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if err := gui.Update(); err != nil {
		return err
	}

	fps.SetText(fmt.Sprintf("%.1f", ebiten.CurrentFPS()))
	frame, err := ebiten.NewImageFromImage(gui.Render(gui.Input.X, gui.Input.Y), ebiten.FilterNearest)
	if err != nil {
		return err
	}
	screen.DrawImage(frame, &ebiten.DrawImageOptions{})
	return nil
}
