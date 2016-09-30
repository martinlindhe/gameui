package ui

import (
	"log"
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/stretchr/testify/assert"
)

func TestEbui(t *testing.T) {

	ui, err := New(320, 240, 2.0)
	if err != nil {
		t.Fatal(err)
	}
	ui.SetWindowTitle("test")

	mainMenu := NewMenuList(ui)
	mainMenu.X = 10
	mainMenu.Y = 10

	newGame := NewMenuItem(ui)
	newGame.Title = "new game"
	newGame.Action = func() {
		log.Println("XXX new game")
	}
	mainMenu.AddItem(newGame)

	cont := NewMenuItem(ui)
	cont.Title = "continue"
	cont.Disabled = true
	mainMenu.AddItem(cont)

	settings := NewMenuItem(ui)
	settings.Title = "settings"
	settings.Disabled = true
	mainMenu.AddItem(settings)

	mnuExit := NewMenuItem(ui)
	mnuExit.Title = "exit"
	mnuExit.Action = func() {
		log.Println("XXX exit")
	}
	mainMenu.AddItem(mnuExit)

	err = ui.AddElement(mainMenu)
	assert.Equal(t, nil, err)

	assert.Equal(t, 4, len(mainMenu.Items))

	//ui.update(im * ebiten.Image)
	err = ui.Loop(updateCallback) // blocking loop
	if err != nil {
		t.Fatal(err)
	}
}

func updateCallback(screen *ebiten.Image) error {

	return nil
}
