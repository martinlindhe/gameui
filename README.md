# About

simple golang lib on top of [image](https://golang.org/pkg/image/) to render a game ui

STATUS: private



# planning

* view contains multiple components:

    window
        - movable
        - closeable (x in corner)
        - minimizable ( _ in corner)
        - resizable
        - click to focus window (put on top of other windows)
        - interact with content (how? child components?)

    button
        - clickable

* each component renders to a separate surface
* blit all surfaces together to produce end UI
* blit end ui on top of game scene to produce end result


