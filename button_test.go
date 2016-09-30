package ui

import (
	"image"
	"image/color"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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
		im, err := btn.Draw()
		assert.Equal(t, nil, err)
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

	im, err := btn.Draw()
	assert.Equal(t, nil, err)
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

	im, err = btn.Draw()
	assert.Equal(t, nil, err)
	testCompareRender(t, []string{
		"#########",
		"#  # #  #",
		"#       #",
		"#   #   #",
		"#########",
	}, renderAsText(im))
}

// for testing
func renderAsText(img *image.RGBA) []string {
	b := img.Bounds()
	res := []string{}
	for y := 0; y < b.Max.Y; y++ {
		row := ""
		for x := 0; x < b.Max.X; x++ {
			col := img.At(x, y)
			row += colToText(col)
		}
		res = append(res, row)
	}
	return res
}

func scale(valueIn, baseMin, baseMax, limitMin, limitMax float64) float64 {
	return ((limitMax - limitMin) * (valueIn - baseMin) / (baseMax - baseMin)) + limitMin
}

// turn col brightness into ascii
func colToText(col color.Color) string {
	vals := []string{
		" ", ".", ",", "+", "o", "5", "6", "O", "0", "#",
	}
	r, g, b, _ := col.RGBA()
	avg := (r + g + b) / 3
	// XXX include alpha by using it as pct of value c
	n := int(scale(float64(avg), 0, 0xffff, 0, 9))
	if n > len(vals) {
		log.Fatal("XXX n too long ", n, len(vals))
	}
	return vals[n]
}

func testCompareRender(t *testing.T, expected, got []string) {
	if len(expected) != len(got) {
		t.Error("expected", len(expected), "lines,got", len(got))
	}
	for i, ex := range expected {
		if i >= len(got) {
			t.Error("line", i+1, "expected", ex, "GOT NOTHING")
			continue
		}
		if ex != got[i] {
			t.Error("line", i+1, "expected", ex, "got", got[i])
		}
	}
}
