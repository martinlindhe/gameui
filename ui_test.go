package ui

import (
	"image"
	"testing"
)

// BenchmarkUI-2   	200000000	         6.07 ns/op
func BenchmarkUI(b *testing.B) {
	w, h := 20, 10
	ui := New(w, h)

	btn := NewButton(w-10, h-4)
	btn.Position = image.Point{X: 5, Y: 3}
	ui.AddComponent(btn)

	txt := NewText("HELLO", 6)
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

	txt := NewText("HELLO", 6)
	txt.Position = image.Point{X: 0, Y: 0}
	ui.AddComponent(txt)

	testCompareRender(t, []string{
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
	}, renderAsText(ui.Render(0, 0)))
}
