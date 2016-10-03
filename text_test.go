package ui

import (
	"fmt"
	"image/color"
	"testing"
)

// BenchmarkDrawText-2     	200000000	         7.36 ns/op   (mbp-2010)
func BenchmarkDrawText(b *testing.B) {

	txt := NewText("HEJ", 6, color.White)
	for n := 0; n < b.N; n++ {
		txt.Draw(0, 0)
	}
}

// BenchmarkDrawChangingText-2   	   10000	    127473 ns/op
func BenchmarkDrawChangingText(b *testing.B) {

	txt := NewText("HEJ", 6, color.White)
	for n := 0; n < b.N; n++ {
		txt.SetText(fmt.Sprintf("hej %d", n))
		txt.Draw(0, 0)
	}
}

func TestTextOnly(t *testing.T) {
	txt := NewText("HEJ", 6, color.White)

	// XXX fixme height is too high
	ex := []string{
		"# # ###   # ",
		"# # ##    # ",
		"### #   # # ",
		"# # ###  #  ",
		"            ",
	}
	// render 2 frames, the second should reach txt.IsClean code paths
	testCompareRender(t, ex, renderAsText(txt.Draw(0, 0)))
	testCompareRender(t, ex, renderAsText(txt.Draw(0, 0)))

	txt.SetText("HOPP")
	ex2 := []string{
		"# #  ##  ##  ##  ",
		"# # #  # # # # # ",
		"### #  # ##  ##  ",
		"# #  ##  #   #   ",
		"                 ",
	}
	testCompareRender(t, ex2, renderAsText(txt.Draw(0, 0))) // XXX
}
