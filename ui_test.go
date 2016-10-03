package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUI(t *testing.T) {
	w, h := 30, 30
	ui := New(w, h)

	btn := NewButton(20, 10)
	btn.Position = Position{X: 5, Y: 5}
	ui.AddComponent(btn)

	txt := NewText("hello", 6)
	txt.Position = Position{X: 0, Y: 0}
	ui.AddComponent(txt)

	assert.Equal(t, 2, len(ui.components))

	// XXX render all components
	// ui.RenderFrame()
}
