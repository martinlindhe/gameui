package ui

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		" ", ".", ",", "3", "4", "5", "6", "7", "8", "9",
	}
	r, g, b, _ := col.RGBA()
	avg := (r + g + b/3)
	// XXX include alpha by using it as pct of value c
	n := int(scale(float64(avg), 0, 255, 0, 9))
	return vals[n]
}

func TestButtonOnly(t *testing.T) {
	w, h := 30, 10
	btn := NewButton(w, h)

	im, err := btn.Draw()
	assert.Equal(t, nil, err)
	txt := renderAsText(im)
	assert.Equal(t, "xxx", txt)

	//ui := New(w, h)
	// ui.AddComponent(btn)
}
