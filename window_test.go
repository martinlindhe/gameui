package ui

import "testing"

func TestWindow(t *testing.T) {
	w, h := 30, 20
	wnd := NewWindow(w, h)

	btn := NewButton(20, 14).SetText("HI")
	btn.Position = Point{X: 5, Y: 3}
	wnd.AddChild(btn)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := wnd.Draw(0, 0)
		testCompareRender(t, []string{
			"##############################",
			"#                            #",
			"#                            #",
			"#    ####################    #",
			"#    #                  #    #",
			"#    #                  #    #",
			"#    # 66. o6+  666666+ #    #",
			"#    # ##. 6#o  ######o #    #",
			"#    # ##. 6#o  oo0#6o, #    #",
			"#    # ##. 6#o    O#,   #    #",
			"#    # ##+,O#o    O#,   #    #",
			"#    # ######o    O#,   #    #",
			"#    # ##OO0#o  ..O#+.  #    #",
			"#    # ##. 6#o  ######o #    #",
			"#    # ##. 6#o  ######o #    #",
			"#    #                  #    #",
			"#    ####################    #",
			"#                            #",
			"#                            #",
			"##############################",
		}, renderAsText(im))
	}
}
