package ui

import (
	"fmt"
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
	fnt, err := NewFont(tinyFontName, 6, 72, White)
	assert.Equal(t, nil, err)

	im, err := fnt.Print("HEJ")
	assert.Equal(t, nil, err)

	testCompareRender(t, []string{
		"# # ###   # ",
		"# # ##    # ",
		"### #   # # ",
		"# # ###  #  ",
		"            ",
		"            ",
		"            ",
		"            ",
		"            ",
		"            ",
	}, renderAsText(im))
}

func TestFontCache(t *testing.T) {
	fnt, err := NewFont(tinyFontName, 6, 72, White)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(fnt.cachedPrints))

	// fill up cache
	for i := 0; i < fontRenderCache; i++ {
		_, err := fnt.Print(fmt.Sprintf("hej %d", i))
		assert.Equal(t, nil, err)
	}

	assert.Equal(t, 10, len(fnt.cachedPrints))
	fnt.Print("hej 999")
	assert.Equal(t, 10, len(fnt.cachedPrints))
}
