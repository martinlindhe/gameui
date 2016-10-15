package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	benchFont, _ = NewFont(defaultFontName, 6, 72, White)
)

// BenchmarkDrawFont-4             100000000               10.7 ns/op (elitebook)
func BenchmarkDrawFont(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchFont.Print("hello world")
	}
}

func TestFontOnly(t *testing.T) {
	fnt, err := NewFont(defaultFontName, 6, 72, White)
	assert.Equal(t, nil, err)

	im, err := fnt.Print("HEJ")

	// imaging.Save(im, "font-only.png")

	assert.Equal(t, nil, err)
	testCompareRender(t, []string{
		"# # ###   # ",
		"# # ##    # ",
		"### #   # # ",
		"# # ###  #  ",
		"            ",
	}, renderAsText(im))
}
