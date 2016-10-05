package ui

import "testing"

func TestWindow(t *testing.T) {
	w, h := 30, 50
	wnd := NewWindow(w, h)

	icongrp := NewIconGroup(16, 16, 16, 16)
	wnd.AddChild(icongrp)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := wnd.Draw(0, 0)
		testCompareRender(t, []string{
			"#########",
			"#       #",
			"#       #",
			"#       #",
			"#########",
		}, renderAsText(im))
	}
}
