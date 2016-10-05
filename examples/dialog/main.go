// a dialog window, with a yes & no button

package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/martinlindhe/farm/ui"
)

const (
	width, height = 320, 200
)

var (
	gui = ui.New(width, height)
	fps = ui.NewText(20, color.White)
)

func main() {
	// XXX improve positions
	exit := ui.NewText(20, color.White).SetText("exit?")
	exit.Position = ui.Point{X: width / 3, Y: height / 3}
	gui.AddComponent(exit)

	btnYes := ui.NewButton(60, 20).SetText("YES")
	btnYes.Position = ui.Point{X: 0, Y: height / 2}
	btnYes.OnClick = func() {
		fmt.Println("clicked", btnYes.Text.GetText())
		os.Exit(0)
	}
	gui.AddComponent(btnYes)

	btnNo := ui.NewButton(60, 20).SetText("NO")
	btnNo.Position = ui.Point{X: width / 2, Y: height / 2}
	btnNo.OnClick = func() {
		fmt.Println("clicked", btnNo.Text.GetText())
	}
	gui.AddComponent(btnNo)

	gui.AddComponent(fps)

	if err := ebiten.Run(update, width, height, 1, "Dialog (UI Demo)"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	gui.Click()

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
