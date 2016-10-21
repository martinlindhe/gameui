package ui

import (
	"image"
	"testing"
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
	w, h := 16, 7
	btn := NewButton(w, h)
	btn.Text.setFont(tinyFontName)
	btn.SetText("HEJ")

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := btn.Draw(0, 0)
		testCompareRender(t, []string{
			"################",
			"# # # ###   #  #",
			"# # # ##    #  #",
			"# ### #   # #  #",
			"# # # ###  #   #",
			"#              #",
			"################",
		}, renderAsText(im))
	}
}

func TestUIWithButtonOnly(t *testing.T) {
	w, h := 9, 5
	ui := New(w, h)

	btn := NewButton(w, h)
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

	r := image.Rect(0, 0, 3, 3)
	im := btn.Draw(0, 0)
	testCompareRender(t, []string{
		"#########",
		"#       #",
		"#       #",
		"#       #",
		"#########",
	}, renderAsText(im))

	icon := image.NewRGBA(r)
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

	icon2 := image.NewRGBA(r)
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
