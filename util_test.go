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

// asserts that expected == got, or fails test
func testCompareRender(t *testing.T, expected, got []string) {
	fail := false
	if len(expected) != len(got) {
		t.Error("expected", len(expected), "lines, got", len(got))
		fail = true
	}
	for i, ex := range expected {
		if i >= len(got) || ex != got[i] {
			fail = true
		}
	}
	if fail {
		for _, g := range got {
			t.Log("_" + g + "_")
		}
		t.FailNow()
	}
}
