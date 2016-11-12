package ui

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	benchButton = NewButton(30, 8)
)

// BenchmarkDrawButton-2   	200000000	         5.58 ns/op    (mbp-2010)
func BenchmarkDrawButton(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchButton.Draw(0, 0)
	}
}

func TestButtonOnly(t *testing.T) {
	w, h := 9, 5
	btn := NewButton(w, h)
	btn.SetBorderColor(White)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := btn.Draw(0, 0)
		testCompareRender(t, []string{
			"#########",
			"#       #",
			"#       #",
			"#       #",
			"#########",
		}, renderAsText(im))
	}
}

func TestButtonWithText(t *testing.T) {
	w, h := 20, 10
	btn := NewButton(w, h)
	btn.SetBorderColor(White)

	tinyFont, err := NewFont(tinyFontName, 7, 72, White)
	assert.Equal(t, nil, err)

	btn.SetText(tinyFont, "HEJ")

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := btn.Draw(0, 0)
		testCompareRender(t, []string{
			"####################",
			"#                  #",
			"#6.o+ 666+   o+    #",
			"##.6o ##6,   6o    #",
			"##+Oo #6.  , 6o    #",
			"##O0o #,.  O,5+    #",
			"##.6o ###o  O,     #",
			"#                  #",
			"#                  #",
			"####################",
		}, renderAsText(im))
	}
}

func TestButtonUIOnlyComponent(t *testing.T) {
	w, h := 9, 5
	ui := New(w, h)

	btn := NewButton(w, h)
	btn.SetBorderColor(White)
	ui.AddComponent(btn)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := ui.Render(0, 0)
		testCompareRender(t, []string{
			"#########",
			"#       #",
			"#       #",
			"#       #",
			"#########",
		}, renderAsText(im))
	}
}

func TestButtonImage(t *testing.T) {
	w, h := 9, 5
	btn := NewButton(w, h)
	btn.SetBorderColor(White)

	r := image.Rect(0, 0, 3, 3)
	im := btn.Draw(0, 0)
	testCompareRender(t, []string{
		"#########",
		"#       #",
		"#       #",
		"#       #",
		"#########",
	}, renderAsText(im))

	icon := image.NewNRGBA(r)
	icon.Set(0, 0, White)
	icon.Set(2, 0, White)
	icon.Set(1, 2, White)
	btn.SetIcon(icon)

	im = btn.Draw(0, 0)
	testCompareRender(t, []string{
		"#########",
		"#  # #  #",
		"#       #",
		"#   #   #",
		"#########",
	}, renderAsText(im))

	icon2 := image.NewNRGBA(r)
	icon2.Set(0, 0, White)
	icon2.Set(1, 0, White)
	icon2.Set(2, 0, White)
	btn.SetIcon(icon2)

	// test render after changed icon
	im = btn.Draw(0, 0)
	testCompareRender(t, []string{
		"#########",
		"#  ###  #",
		"#       #",
		"#       #",
		"#########",
	}, renderAsText(im))
}
