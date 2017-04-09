package ui

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myItem struct {
	text string
}

func (o myItem) Text() string {
	return o.text
}

func TestListOnly(t *testing.T) {

	grp := NewList(70, 30)

	fnt, err := NewFont(defaultFontName, 12, 72, color.White)
	assert.Equal(t, nil, err)

	o1 := myItem{text: "ITEM 1"}
	grp.AddLine(o1, fnt, func() {})

	o2 := myItem{text: "ITEM 2"}
	grp.AddLine(o2, fnt, func() {})

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := grp.Draw(-1, -1)
		testCompareRender(t, []string{
			"                                                                      ",
			"     ,,,,,,,,.          .,.    ,.                                     ",
			" 5#. OO6O#66Oo  0####0  5#6   5#o            .+o6                     ",
			" +#     o0      #       5#0   0#5            o0O0                     ",
			" .0     o#      #       5o#+ ,#o5               0                     ",
			" .0     5#      #       6.#6 6#.6               0                     ",
			" .0     5#      #####O  6 O0.0O 6               0                     ",
			" ,0     5#.     #       6 o#O#o 6               0                     ",
			" +0     6#.     #       O .###. O               0                     ",
			" 5#     6#.     #       O  O#O  O              .0                     ",
			" O#.    6#,     ######, 0  ,o,  0              .0                     ",
			" 0#+    6#,     0#####, 0       0             0###5                   ",
			"                                                                      ",
			"     ,,,,,,,,.          .,.    ,.            .500O,                   ",
			" 5#. OO6O#66Oo  0####0  5#6   5#o            o5. +O                   ",
			" +#     o0      #       5#0   0#5                 O                   ",
			" .0     o#      #       5o#+ ,#o5                 O                   ",
			" .0     5#      #       6.#6 6#.6                ,5                   ",
			" .0     5#      #####O  6 O0.0O 6               .0.                   ",
			" ,0     5#.     #       6 o#O#o 6              .0o                    ",
			" +0     6#.     #       O .###. O             .06                     ",
			" 5#     6#.     #       O  O#O  O            ,0O                      ",
			" O#.    6#,     ######, 0  ,o,  0            5####0                   ",
			" 0#+    6#,     0#####, 0       0            5####0                   ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
		}, renderAsText(im))
	}
}
