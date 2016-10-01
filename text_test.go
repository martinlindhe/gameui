package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BenchmarkDrawText-2     	200000000	         7.15 ns/op   (mbp-2010)
func BenchmarkDrawText(b *testing.B) {

	txt := NewText("hej", 6)
	for n := 0; n < b.N; n++ {
		txt.Draw()
	}
}

func TestTextOnly(t *testing.T) {
	txt := NewText("hej", 6)

	// XXX fixme height is too high
	im, err := txt.Draw()
	assert.Equal(t, nil, err)
	testCompareRender(t, []string{
		"# # ###   # ",
		"# # ##    # ",
		"### #   # # ",
		"# # ###  #  ",
		"            ",
	}, renderAsText(im))
}
