package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkUI-4                   500000000                3.52 ns/op (elitebook)
func BenchmarkUI(b *testing.B) {
	w, h := 20, 10
	ui := New(w, h)

	btn := NewButton(w-10, h-4)
	btn.Position = Point{X: 5, Y: 3}
	ui.AddComponent(btn)

	txt := NewText(6, White)
	txt.SetText("HELLO")
	txt.Position = Point{X: 0, Y: 0}
	ui.AddComponent(txt)

	for n := 0; n < b.N; n++ {
		ui.Render(0, 0)
	}
}

func TestUI(t *testing.T) {
	w, h := 20, 10
	ui := New(w, h)
	ui.SetWindowTitle("test ui")

	btn := NewButton(w-10, h-4)
	btn.Position = Point{X: 5, Y: 3}
	ui.AddComponent(btn)

	txt := NewText(6, White)
	txt.setFont(tinyFontName)
	txt.SetText("HELLO")
	txt.Show()
	assert.Equal(t, false, txt.IsHidden())
	ui.AddComponent(txt)

	hidden := NewText(6, White)
	hidden.SetText("INVISIBLE")
	hidden.Hide()
	assert.Equal(t, true, hidden.IsHidden())
	ui.AddComponent(hidden)

	assert.Equal(t, true, CheckUI(ui))

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

func TestUIClick(t *testing.T) {
	w, h := 20, 20
	ui := New(w, h)
	ui.SetWindowTitle("test ui")

	// create a window, with a button in it
	wnd := NewWindow(w, h)
	ui.AddComponent(wnd)

	clicked := false
	btn := NewButton(w-10, h-4)
	btn.Position = Point{X: 5, Y: 3}
	btn.OnClick = func() {
		clicked = true
	}
	wnd.AddChild(btn)

	ui.handleClick() // try without setting mouse button to reach other code path

	// fake left mouse click in the area of button
	ui.Input.mouseStates[MouseButtonLeft] = 1
	ui.Input.X = btn.Position.X + 1 + 2
	ui.Input.Y = btn.Position.Y + 1 + 12 + 2 // 12 = window title
	ui.handleClick()
	assert.Equal(t, true, clicked)
}
