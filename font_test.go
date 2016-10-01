package ui

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	benchFont, _ = NewFont("../_resources/font/tiny/tiny.ttf", 6, 72, color.White)
)

// BenchmarkDrawFont-2     	  100000	     14122 ns/op   (mbp-2010)
func BenchmarkDrawFont(b *testing.B) {

	for n := 0; n < b.N; n++ {
		benchFont.Print("hello world")
	}
}

func TestFontOnly(t *testing.T) {
	fnt, err := NewFont("../_resources/font/tiny/tiny.ttf", 6, 72, color.White)
	assert.Equal(t, nil, err)

	// XXX Print should only return as big image as needed. height is wrong!
	im, err := fnt.Print("HEJ")
	assert.Equal(t, nil, err)
	testCompareRender(t, []string{
		"# # ###   # ",
		"# # ##    # ",
		"### #   # # ",
		"# # ###  #  ",
		"            ",
	}, renderAsText(im))
}
