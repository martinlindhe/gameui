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

	grp := NewList(100, 50)

	o1 := babj{name: "item1"}
	grp.AddLine(o1)

	o2 := babj{name: "item2"}
	grp.AddLine(o2)

	o3 := babj{name: "item3"}
	grp.AddLine(o3)

	// make sure same frame is delivered each time
	for i := 0; i < 10; i++ {
		im := grp.Draw(-1, -1)
		testCompareRender(t, []string{
			"##########",
			"#        #",
			"# # ##   #",
			"#     #  #",
			"#  #   # #",
			"# # #    #",
			"#        #",
			"#  #     #",
			"#        #",
			"##########",
		}, renderAsText(im))
	}
}
