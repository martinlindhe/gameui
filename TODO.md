# TODO

* window
    - movable
    - closeable (x in corner)
    - resizable
    - click to focus window (put on top of other windows)

* later move "ui" to ebui pkg, overwriting existing project there

* no ebiten dependency! only export an image
    !!! currently the input.go binds to ebiten input reading

* DEMO: rework the ebui menu scroll thingy
    vert list of buttons with text

* font: render all letters in one line to a image.Image ONCE per used size, when printing, just blit from parts
    of this source image
