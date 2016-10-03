package ui

import (
	"image"
	"testing"
)

func TestUI(t *testing.T) {
	w, h := 20, 10
	ui := New(w, h)

	btn := NewButton(w-10, h-4)
	btn.Position = image.Point{X: 5, Y: 4}
	ui.AddComponent(btn)

	txt := NewText("HELLO", 6)
	txt.Position = image.Point{X: 0, Y: 0}
	ui.AddComponent(txt)

	// XXX with btn.Y = 4, it should overlap text!!!
	testCompareRender(t, []string{
		"# # ### #   #    ## ",
		"# # ##  #   #   #  #",
		"### #   #   #   #  #",
		"# # ### ### ###  ## ",
		"     ##########     ",
		"     #        #     ",
		"     #        #     ",
		"     #        #     ",
		"     #        #     ",
		"     ##########     ",
	}, renderAsText(ui.Render()))
}
