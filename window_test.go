package ui

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindow(t *testing.T) {
	w, h := 30, 20
	wnd := NewWindow(w, h, "")
	wnd.SetTitleColor(color.Black)
	wnd.SetBackgroundColor(color.Black)
	wnd.SetBorderColor(color.White)
	wnd.HideCloseButton()

	tinyFont, err := NewFont(tinyFontName, 11, 72, White)
	assert.Equal(t, nil, err)

	btn := NewButton(20, 14).SetText(tinyFont, "HI")
	btn.SetBorderColor(White)
	btn.Position = Point{X: 5, Y: 3}
	wnd.AddChild(btn)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := wnd.Draw(0, 0)
		testCompareRender(t, []string{
			"##############################",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,####################,,,,#",
			"#,,,,#,,,,,,,,,,,,,,,,,,#,,,,#",
			"#,,,,#,,,,,,,,,,,,,,,,,,#,,,,#",
			"#,,,,#oo,,o+,ooooo+,,,,,#,,,,#",
			"#,,,,##O,o#5,#####5,,,,,#,,,,#",
			"#,,,,##O,o#5,56#O5+,,,,,#,,,,#",
			"#,,,,##O,o#5,,+#6,,,,,,,#,,,,#",
			"#,,,,##06O#5,,+#6,,,,,,,#,,,,#",
			"#,,,,######5,,+#6,,,,,,,#,,,,#",
			"#,,,,##0+5#5,O0#0O5,,,,,#,,,,#",
			"#,,,,##O,o#5,#####5,,,,,#,,,,#",
			"#,,,,#,,,,,,,,,,,,,,,,,,#,,,,#",
			"#,,,,#,,,,,,,,,,,,,,,,,,#,,,,#",
			"#,,,,####################,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"##############################",
		}, renderAsText(im))
	}
}

func TestWindowWithTitle(t *testing.T) {
	w, h := 40, 20
	wnd := NewWindow(w, h, "WOA")
	wnd.SetTitleColor(color.Black)
	wnd.SetBackgroundColor(color.Black)
	wnd.SetBorderColor(White)
	wnd.HideCloseButton()

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := wnd.Draw(0, 0)
		testCompareRender(t, []string{
			"########################################",
			"#+5,,,,,,,,oo,,+O0Oo,,,,,,6+,,,,,,,,,,,#",
			"#,0+,,,,,,,0o,+O,,,Oo,,,,5#5,,,,,,,,,,,#",
			"#,O5,,,,,,o0,,6+,,,+0,,,,O60,,,,,,,,,,,#",
			"#,50,,+o,,OO,,O,,,,,0+,,+0,0o,,,,,,,,,,#",
			"#,+#+,00,+05,,O,,,,,O+,,66,6O,,,,,,,,,,#",
			"#,,065##55#+,,O,,,,,0+,,0o,+0+,,,,,,,,,#",
			"#,,O000O00O,,,65,,,50,,o#####5,,,,,,,,,#",
			"#,,5##5o##5,,,+0####o,,O6,,,50,,,,,,,,,#",
			"#,,+00,,O0+,,,,+O0Oo,,+#5,,,o#o,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"#,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,#",
			"########################################",
		}, renderAsText(im))
	}
}
