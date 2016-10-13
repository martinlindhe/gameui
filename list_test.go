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

	o3 := babj{name: "ITEM 3", color: color.White}
	grp.AddLine(o3, func() {})

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
			"                                                                      ",
			"######  ######  ######  ##      ##          ##                        ",
			"######  ######  ######  ##      ##          ##                        ",
			"  ##      ##    ####    ####  ####        ####                        ",
			"  ##      ##    ####    ####  ####        ####                        ",
			"  ##      ##    ##      ##  ##  ##          ##                        ",
			"  ##      ##    ##      ##  ##  ##          ##                        ",
			"######    ##    ######  ##      ##        ######                      ",
			"######    ##    ######  ##      ##        ######                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"######  ######  ######  ##      ##        ####                        ",
			"######  ######  ######  ##      ##        ####                        ",
			"  ##      ##    ####    ####  ####            ##                      ",
			"  ##      ##    ####    ####  ####            ##                      ",
			"  ##      ##    ##      ##  ##  ##        ##                          ",
			"  ##      ##    ##      ##  ##  ##        ##                          ",
			"######    ##    ######  ##      ##        ######                      ",
			"######    ##    ######  ##      ##        ######                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"######  ######  ######  ##      ##        ####                        ",
			"######  ######  ######  ##      ##        ####                        ",
			"  ##      ##    ####    ####  ####          ####                      ",
			"  ##      ##    ####    ####  ####          ####                      ",
			"  ##      ##    ##      ##  ##  ##            ##                      ",
			"  ##      ##    ##      ##  ##  ##            ##                      ",
			"######    ##    ######  ##      ##        ####                        ",
			"######    ##    ######  ##      ##        ####                        ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
			"                                                                      ",
		}, renderAsText(im))
	}
}
