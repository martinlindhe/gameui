# About

[![GoDoc](https://godoc.org/github.com/martinlindhe/gameui?status.svg)](https://godoc.org/github.com/martinlindhe/gameui)
[![Travis-CI](https://api.travis-ci.org/martinlindhe/gameui.svg)](https://travis-ci.org/martinlindhe/gameui)
[![codecov.io](https://codecov.io/github/martinlindhe/gameui/coverage.svg?branch=master)](https://codecov.io/github/martinlindhe/gameui?branch=master)

A game UI for golang with boilerplate to run with [ebiten](https://github.com/hajimehoshi/ebiten/).


# Example usage

```go
package main

import "github.com/martinlindhe/gameui"

const (
    width, height = 320, 200
    fontName      = "_resources/font/open_dyslexic/OpenDyslexic3-Regular.ttf"
)

var (
    gui       = ui.New(width, height)
    font12, _ = ui.NewFont(fontName, 12, 72, ui.White)
)

func main() {
    // add a text element
    text := ui.NewText(font12).SetText("hello")
    text.Position = ui.Point{X: width/2 - text.GetWidth()/2, Y: height / 3}
    gui.AddComponent(text)

    if err := gui.Update(); err != nil {
        return err
    }

    // get a image.Image
    img := gui.Render()
}
```


Full example using ebiten:

    go get github.com/martinlindhe/gameui
    cd $GOPATH/src/github.com/martinlindhe/gameui
    go run examples/tooltip/main.go

See the [examples folder](examples) for more examples


### License

Under [MIT](LICENSE)
