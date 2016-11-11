package ui

import (
	"image"
	"testing"
)

func TestBarOnly(t *testing.T) {
	w, h := 32, 5
	bar := NewBar(w, h)
	bar.SetValue(50)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := bar.Draw(0, 0)
		testCompareRender(t, []string{
			"++++++++++++++++++++++++++++++++",
			"+666666666666666               +",
			"+666666666666666               +",
			"+666666666666666               +",
			"++++++++++++++++++++++++++++++++",
		}, renderAsText(im))
	}
}

func TestBarWithImageOnly(t *testing.T) {
	w, h := 30, 5
	bar := NewBar(w, h)
	bar.SetValue(50)

	im1 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	im1.Set(0, 0, White)
	im1.Set(2, 0, White)
	im1.Set(1, 2, White)

	bar.SetFillImage(im1)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := bar.Draw(0, 0)
		testCompareRender(t, []string{
			// does not draw the rightmost # (1/3 of last pic) because it overflows max val
			"++++++++++++++++++++++++++++++",
			"+# ## ## ## ##               +",
			"+                            +",
			"+ #  #  #  #  #              +",
			"++++++++++++++++++++++++++++++",
		}, renderAsText(im))
	}
}
