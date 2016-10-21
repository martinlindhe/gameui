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
			"    6    .6666666  ,66666  .6o   56            o66                    ",
			"   .0     ..+#...  +O,,,,  ,#0  .0#           .5o#                    ",
			"   .0       ,#     +6      ,06, o6#              #                    ",
			"   .0       ,#     +O....  ,#+6 O+#              #                    ",
			"   .0       ,#     +0OOO6  +#.0,0,#.             #                    ",
			"   .0       +#.    +6      +# 006,#.             #                    ",
			"   ,#       +#.    +6      +# 6#+,#.             #                    ",
			"   ,#       o#.    +0OOOO. +#.,6.+#.           56#65                  ",
			"   ,0       o#,    +#####. +#.   +#.           0###0                  ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"    6    .6666666  ,66666  .6o   56            5OO6.                  ",
			"   .0     ..+#...  +O,,,,  ,#0  .0#           .o..5O                  ",
			"   .0       ,#     +6      ,06, o6#                #                  ",
			"   .0       ,#     +O....  ,#+6 O+#               ,0                  ",
			"   .0       ,#     +0OOO6  +#.0,0,#.             .0+                  ",
			"   .0       +#.    +6      +# 006,#.            .0o                   ",
			"   ,#       +#.    +6      +# 6#+,#.           +05                    ",
			"   ,#       o#.    +0OOOO. +#.,6.+#.          ,##O66                  ",
			"   ,0       o#,    +#####. +#.   +#.          +#####.                 ",
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
			"                                                                      ",
			"                                                                      ",
		}, renderAsText(im))
	}
}
