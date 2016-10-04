package ui

import (
	"image"
	"image/color"
	"testing"
)

// BenchmarkUI-4                   500000000                3.52 ns/op (elitebook)
func BenchmarkUI(b *testing.B) {
	w, h := 20, 10
	ui := New(w, h)

	btn := NewButton(w-10, h-4)
	btn.Position = image.Point{X: 5, Y: 3}
	ui.AddComponent(btn)

	txt := NewText(6, color.White)
	txt.SetText("HELLO")
	txt.Position = image.Point{X: 0, Y: 0}
	ui.AddComponent(txt)

	for n := 0; n < b.N; n++ {
		ui.Render(0, 0)
	}
}

func TestUI(t *testing.T) {
	w, h := 20, 10
	ui := New(w, h)

	btn := NewButton(w-10, h-4)
	btn.Position = image.Point{X: 5, Y: 3}
	ui.AddComponent(btn)

	txt := NewText(6, color.White)
	txt.SetText("HELLO")
	txt.Position = image.Point{X: 0, Y: 0}
	ui.AddComponent(txt)

	ex := []string{
		"# # ### #   #    ## ",
		"# # ##  #   #   #  #",
		"### #   #   #   #  #",
		"# # ###########  ## ",
		"     #        #     ",
		"     #        #     ",
		"     #        #     ",
		"     #        #     ",
		"     ##########     ",
		"                    ",
	}
	testCompareRender(t, ex, renderAsText(ui.Render(0, 0)))
	testCompareRender(t, ex, renderAsText(ui.Render(0, 0)))

	txt.SetText("BEEP")
	ex2 := []string{
		"##  ### ### ##      ",
		"### ##  ##  # #     ",
		"# # #   #   ##      ",
		"##  ###########     ",
		"     #        #     ",
		"     #        #     ",
		"     #        #     ",
		"     #        #     ",
		"     ##########     ",
		"                    ",
	}
	testCompareRender(t, ex2, renderAsText(ui.Render(0, 0)))
}
