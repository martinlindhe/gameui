package main

import (
	"fmt"
	_ "image/png"
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
	hp      *ui.Bar
	hp2     *ui.Bar
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

	heart, err := ui.OpenImage("_resources/tile/7x7_heart.png")
	if err != nil {
		log.Fatal(err)
	}

	hp = ui.NewBar(width-2, 7+2)
	hp.SetValue(0)
	hp.SetFillImage(heart)
	hp.Position = ui.Point{X: width/2 - hp.Dimension.Width/2, Y: 20}
	gui.AddComponent(hp)

	hp2 = ui.NewBar(width-2, 7)
	hp2.SetValue(0)
	hp2.SetFillColor(ui.Red)
	hp2.Position = ui.Point{X: width/2 - hp.Dimension.Width/2, Y: 30}
	gui.AddComponent(hp2)

	mana = ui.NewBar(width-2, 16)
	mana.SetValue(25)
	mana.SetFillColor(ui.Blue)
	mana.Position = ui.Point{X: width/2 - hp.Dimension.Width/2, Y: height - 16 - 20}
	gui.AddComponent(mana)

	lastInc = time.Now()
}

func main() {
	if err := ebiten.Run(update, width, height, scale, "Bars (UI Demo)"); err != nil {
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

		hp.IncValue(1)
		hp2.IncValue(1)
		mana.IncValue(2)
	}

	frame, err := ebiten.NewImageFromImage(gui.Render(0, 0), ebiten.FilterNearest)
	if err != nil {
		return err
	}
	screen.DrawImage(frame, &ebiten.DrawImageOptions{})
	return nil
}
