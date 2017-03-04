// press space to toggle hide / show grouop of child components

package main

import (
	"fmt"
	"log"

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
	fps       = ui.NewText(font12)
)

func main() {

	grp := ui.NewGroup(210, 16)
	grp.Position = ui.Point{
		X: (width / 2) - (grp.Dimension.Width / 2),
		Y: (height / 2) - (grp.Dimension.Height / 2)}
	gui.AddComponent(grp)

	txt := ui.NewText(font12).SetText("press space to toggle visible")
	txt.Position = ui.Point{X: 0, Y: 0}
	grp.AddChild(txt)
	gui.AddComponent(fps)

	gui.AddKeyFunc(ui.KeySpace, func() error {
		if grp.IsHidden() {
			grp.Show()
		} else {
			grp.Hide()
		}
		return nil
	})

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
	frame, err := ebiten.NewImageFromImage(gui.Render(0, 0), ebiten.FilterNearest)
	if err != nil {
		return err
	}
	screen.DrawImage(frame, &ebiten.DrawImageOptions{})
	return nil
}
