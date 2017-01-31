package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	w, h := 30, 20
	grp := NewGroup(w, h)

	tinyFont, err := NewFont(tinyFontName, 11, 72, White)
	assert.Equal(t, nil, err)

	btn := NewButton(20, 14).SetText(tinyFont, "HI")
	btn.SetBorderColor(White)
	btn.Position = Point{X: 5, Y: 3}
	grp.AddChild(btn)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := grp.Draw(0, 0)
		testCompareRender(t, []string{
			"                              ",
			"                              ",
			"                              ",
			"     ####################     ",
			"     #                  #     ",
			"     #                  #     ",
			"     #,,  ,. ,,,,,.     #     ",
			"     ##O ,#o #####o     #     ",
			"     ##O ,#o o5#Oo,     #     ",
			"     ##O ,#o  .#6       #     ",
			"     ##06O#o  .#6       #     ",
			"     ######o  .#6       #     ",
			"     ##O.+#o OO#0O+     #     ",
			"     ##O ,#o #####o     #     ",
			"     #                  #     ",
			"     #                  #     ",
			"     ####################     ",
			"                              ",
			"                              ",
			"                              ",
		}, renderAsText(im))
	}
}
