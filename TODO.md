# TODO

xxx
* copy stuff from ebui. drop ebui usage atm. later port "ui" to ebui pkg

* no ebiten dependency ?


# planning

xxx

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
* fast text rendering, with unicode
* load ttf font

* test rendering by pixel compare reference images
* dump render result to disk, using the image lib github.com/disintegration/imaging

* speed: benchmark time to render 100 frames
* speed: view tracks dirtiness of components (if they need to be redrawn)
