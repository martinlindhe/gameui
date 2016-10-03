package ui

import "testing"

// BenchmarkDrawText-2     	200000000	         7.36 ns/op   (mbp-2010)
func BenchmarkDrawText(b *testing.B) {

	txt := NewText("HEJ", 6)
	for n := 0; n < b.N; n++ {
		txt.Draw(0, 0)
	}
}

func TestTextOnly(t *testing.T) {
	txt := NewText("HEJ", 6)

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
}
