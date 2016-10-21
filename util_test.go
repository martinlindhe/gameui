package ui

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"testing"

	"github.com/martinlindhe/go-difflib/difflib"
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
	r, g, b, a := col.RGBA()
	avg := (r + g + b + a) / 4

	n := int(scale(float64(avg), 0, 0xffff, 0, 9))
	if n > len(vals) {
		log.Fatal("n is too big", n, len(vals))
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
			break
		}
	}
	if fail {
		diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
			A:        expected,
			B:        got,
			FromFile: "expected",
			ToFile:   "got",
			Context:  3,
			Eol:      "\n",
		})
		fmt.Print(diff)

		t.FailNow()
	}
}
