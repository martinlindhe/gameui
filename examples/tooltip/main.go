package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	ui "github.com/martinlindhe/gameui"
)

const (
	width, height = 320, 200
	scale         = 2.
	fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
	gui     *ui.UI
	fps     *ui.Text
	mana    *ui.Bar
	lastInc time.Time
)

func init() {
	gui = ui.New(width, height)
	fnt, err := ui.NewFont(fontName, 12, 72, ui.White)
	if err != nil {
		log.Fatal(err)
	}
	fps = ui.NewText(fnt)
	gui.AddComponent(fps)
	gui.AddKeyFunc(ui.KeyQ, func() error {
		fmt.Println("q - QUITTING")
		return ui.GracefulExitError{}
	})

	mana = ui.NewBar(width-2, 16)
	mana.SetValue(25)
	mana.SetFillColor(ui.Blue)
	mana.Position = ui.Point{X: 0, Y: (height / 2) + (height / 4)}
	mana.SetTooltip(fmt.Sprintf("mana = %d", mana.GetValue()))
	gui.AddComponent(mana)

	grpWidth := 40
	grp := ui.NewGroup(grpWidth, 100)
	grp.Position = ui.Point{X: (width / 2) - (grpWidth / 2), Y: 10}
	bar1 := ui.NewBar(grpWidth, 10)
	bar1.Position = ui.Point{X: 0, Y: 0}
	bar1.SetValue(10)
	bar1.SetFillColor(ui.Red)
	bar1.SetTooltip("bar 1")
	grp.AddChild(bar1)

	bar2 := ui.NewBar(grpWidth, 10)
	bar2.Position = ui.Point{X: 0, Y: 15}
	bar2.SetValue(20)
	bar2.SetFillColor(ui.Green)
	bar2.SetTooltip("bar 2")
	grp.AddChild(bar2)

	bar3 := ui.NewBar(grpWidth, 10)
	bar3.Position = ui.Point{X: 0, Y: 30}
	bar3.SetValue(40)
	bar3.SetFillColor(ui.Yellow)
	bar3.SetTooltip("bar 3")
	grp.AddChild(bar3)

	gui.AddComponent(grp)

	lastInc = time.Now()
}

func main() {
	if err := ebiten.Run(update, width, height, scale, "Tooltip (UI Demo)"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	if err := gui.Update(); err != nil {
		return err
	}

	fps.SetText(fmt.Sprintf("%.1f", ebiten.CurrentFPS()))

	expired := time.Now().Add(-1 * time.Second)
	if lastInc.Before(expired) {
		lastInc = time.Now()
		mana.IncValue(2)
		mana.SetTooltip(fmt.Sprintf("mana = %d", mana.GetValue()))
	}

	frame, err := ebiten.NewImageFromImage(gui.Render(gui.Input.X, gui.Input.Y), ebiten.FilterNearest)
	if err != nil {
		return err
	}
	screen.DrawImage(frame, &ebiten.DrawImageOptions{})
	return nil
}
