# TODO

* tests
    - raise "make cover":
        ok  	github.com/martinlindhe/gameui	0.312s	coverage: 66.6% of statements

        TEST mouse clicks to reach most remaining code paths for coverage

* Window
    - movable (drag title bar)
    - resizable (drag resize area, bottom right triangle)
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


# Interfaces
Tested (using image.Image and draw.Image instead of `*image.RGBA`) and
saw 4x worse performance doing this (go 1.7.3 darwin)

    before:
    BenchmarkDrawButton-2         200000000	         6.04 ns/op
    BenchmarkDrawFont-2           100000000	        20.6 ns/op
    BenchmarkDrawText-2           300000000	         5.49 ns/op
    BenchmarkDrawChangingText-2   3000000	       463 ns/op
    BenchmarkUI-2                 50000000	        32.3 ns/op

    after:
    BenchmarkDrawButton-2         50000000	        33.2 ns/op
    BenchmarkDrawFont-2           50000000	        22.6 ns/op
    BenchmarkDrawText-2           50000000	        31.6 ns/op
    BenchmarkDrawChangingText-2   3000000	       471 ns/op
    BenchmarkUI-2                 30000000	        53.8 ns/op
