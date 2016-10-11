package ui

import (
	"image"
	"testing"
)

// obj implements IconGroupObject interface
type obj struct {
	name string
	icon *image.RGBA
	id   uint64
}

func (o obj) Name() string {
	return o.name
}
func (o obj) Icon() *image.RGBA {
	return o.icon
}
func (o obj) ID() uint64 {
	return o.id
}
func (o obj) Click() {
	// XXX fmt.Println("CLICKED ME: ", o.name)
}

func TestIcongroupOnly(t *testing.T) {

	im1 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	im1.Set(0, 0, White)
	im1.Set(2, 0, White)
	im1.Set(1, 2, White)

	im2 := image.NewRGBA(image.Rect(0, 0, 3, 3))
	im2.Set(0, 0, White)
	im2.Set(1, 1, White)
	im2.Set(2, 2, White)

	grp := NewIconGroup(2, 2, 3, 3)

	o1 := obj{name: "item1", icon: im1, id: 1}
	grp.AddObject(o1)

	o2 := obj{name: "item2", icon: im2, id: 2}
	grp.AddObject(o2)

	// make sure second row of objects is shown in ui
	o3 := obj{name: "item3", icon: im1, id: 3}
	grp.AddObject(o3)

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
