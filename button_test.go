package ui

import (
	"image"
	"image/color"
	"testing"
)

var (
	benchButton = NewButton(30, 8)
)

// BenchmarkDrawButton-2   	200000000	         6.53 ns/op    (mbp-2010)
func BenchmarkDrawButton(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchButton.Draw()
	}
}

func TestButtonOnly(t *testing.T) {
	w, h := 9, 5
	btn := NewButton(w, h)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := btn.Draw()
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

	r := image.Rect(0, 0, 3, 3)
	img := image.NewRGBA(r)

	im := btn.Draw()
	testCompareRender(t, []string{
		"#########",
		"#       #",
		"#       #",
		"#       #",
		"#########",
	}, renderAsText(im))

	img.Set(0, 0, color.White)
	img.Set(2, 0, color.White)
	img.Set(1, 2, color.White)

	btn.SetImage(img)

	im = btn.Draw()
	testCompareRender(t, []string{
		"#########",
		"#  # #  #",
		"#       #",
		"#   #   #",
		"#########",
	}, renderAsText(im))
}
