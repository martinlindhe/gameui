# TODO

* Window
    - movable
    - resizable
    - click to focus window (put on top of other windows)

* IconGroup
    - allow scrolling with scroll bar / mouse wheel
    - show o.Name() if mouse is hover without click

* Font
    - create a one-row wide lookup image with all letters rendered? (bitmap font mode)
    - google/freetype is very slow, have a look at https://github.com/google/font-go

* LATER
    - move "ui" to ebui pkg, overwriting existing project there

* no ebiten dependency! only export an image
    !!! currently the input.go binds to ebiten input reading

* DEMO: rework the ebui menu scroll thingy
    vert list of buttons with text

* font: render all letters in one line to a image.Image ONCE per used size, when printing, just blit from parts
    of this source image
