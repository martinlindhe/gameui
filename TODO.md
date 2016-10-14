# TODO

* tests
    - raise "make cover":
        ok  	github.com/martinlindhe/farm/ui	0.312s	coverage: 66.3% of statements

        TEST mouse clicks to reach most remaining code paths for coverage

* Window
    - movable
    - resizable
    - click to focus window (put on top of other windows)

* IconGroup
    - allow scrolling with scroll bar / mouse wheel
    - show o.Name() if mouse is hover without click

* Font
    - create a one-row wide lookup image with all letters rendered? (bitmap font mode)
    - render all letters in one line to a image.Image ONCE per used size, when printing,
        just blit from parts of this source image

    - google/freetype is very slow, have a look at https://github.com/google/font-go

* LATER
    - move "ui" to ebui pkg, overwriting existing project there

* dependencies
    - ebiten: no ebiten dependency! only export an image
        currently input.go binds to ebiten input reading

* examples
    - rework the ebui menu scroll thing. vert list of buttons with text
