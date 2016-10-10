package ui

import (
	"image"
	"image/color"
	"log"
	"testing"
)

func scale(valueIn, baseMin, baseMax, limitMin, limitMax float64) float64 {
	return ((limitMax - limitMin) * (valueIn - baseMin) / (baseMax - baseMin)) + limitMin
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
		t.Error("expected", len(expected), "lines, got", len(got))
		for _, g := range got {
			t.Log(g)
		}
		t.FailNow()
		return
	}
	for i, ex := range expected {
		if i >= len(got) {
			t.Error("line", i+1, "expected", ex, "GOT NOTHING")
			continue
		}
		if ex != got[i] {
			t.Error("line", i+1, "expected", ex, "got", got[i], ".")
		}
	}
}
