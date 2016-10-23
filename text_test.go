package ui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkDrawText-2     	200000000	         7.36 ns/op   (mbp-2010)
func BenchmarkDrawText(b *testing.B) {
	txt := NewText(6, White)
	txt.SetText("HEJ")
	for n := 0; n < b.N; n++ {
		txt.Draw(0, 0)
	}
}

// BenchmarkDrawChangingText-4     10000000               262 ns/op (elitebook)
func BenchmarkDrawChangingText(b *testing.B) {
	txt := NewText(6, White)
	for n := 0; n < b.N; n++ {
		s := fmt.Sprintf("hej %d", n%8)
		txt.SetText(s)
		txt.Draw(0, 0)
	}
}

func TestTextOnly(t *testing.T) {
	txt := NewText(6, White)

	ex := []string{
		"            ",
		"            ",
		"# # ###   # ",
		"# # ##    # ",
		"### #   # # ",
		"# # ###  #  ",
		"            ",
		"            ",
		"            ",
		"            ",
	}
	txt.setFont(tinyFontName)
	txt.SetText("HEJ")
	assert.Equal(t, "HEJ", txt.GetText())
	assert.Equal(t, 12, txt.GetWidth())

	// render 2 frames, the second should reach txt.IsClean code paths
	testCompareRender(t, ex, renderAsText(txt.Draw(0, 0)))
	testCompareRender(t, ex, renderAsText(txt.Draw(0, 0)))

	// change text, make sure the change is rendered
	txt.SetText("HOPP")
	ex2 := []string{
		"                 ",
		"                 ",
		"# #  ##  ##  ##  ",
		"# # #  # # # # # ",
		"### #  # ##  ##  ",
		"# #  ##  #   #   ",
		"                 ",
		"                 ",
		"                 ",
		"                 ",
	}
	testCompareRender(t, ex2, renderAsText(txt.Draw(0, 0)))
}
