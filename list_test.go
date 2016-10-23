package ui

import (
	"image/color"
	"testing"
)

type babj struct {
	name  string
	color color.Color
}

func (o babj) Name() string {
	return o.name
}
func (o babj) Color() color.Color {
	return o.color
}
func TestListOnly(t *testing.T) {

	grp := NewList(70, 50)

	o1 := babj{name: "ITEM 1", color: color.White}
	grp.AddLine(o1, func() {})

	o2 := babj{name: "ITEM 2", color: color.White}
	grp.AddLine(o2, func() {})

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := grp.Draw(-1, -1)
		testCompareRender(t, []string{
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
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
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
		}, renderAsText(im))
	}
}
