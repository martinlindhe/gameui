package ui

import "image"

// Group is a container of more components, similar to a transparent Window
type Group struct {
	component
}

// NewGroup creates a new Group
func NewGroup(width, height int) *Group {
	grp := Group{}
	grp.Dimension.Width = width
	grp.Dimension.Height = height
	return &grp
}

// AddChild adds a child component to the Group
func (grp *Group) AddChild(c Component) {
	grp.isClean = false
	grp.children = append(grp.children, c)
}

// Draw redraws internal buffer
func (grp *Group) Draw(mx, my int) *image.RGBA {
	if grp.isHidden {
		return nil
	}
	if grp.isClean {
		if grp.isChildrenClean() {
			return grp.Image
		}
		grp.isClean = false
	}
	grp.initImage()

	grp.drawChildren(mx, my)

	grp.isClean = true
	return grp.Image
}

// Click pass click to window child components
func (grp *Group) Click(mouse Point) bool {
	childPoint := Point{X: mouse.X - grp.Position.X, Y: mouse.Y - grp.Position.Y}
	for _, c := range grp.children {
		if childPoint.In(c.GetBounds()) {
			c.Click(childPoint)
			return true
		}
	}
	return false
}
