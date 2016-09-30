# TODO

* later move "ui" to ebui pkg, overwriting existing project there

* no ebiten dependency ? only export an image


* test: render component, check result -> output rough ascii img for debug purposes
    map color brightness to symbols "space" to # (little-to-much-covered box) A-Z and then render to string, use in tests


* font: render all letters in one line to a image.Image ONCE, when printing, just blit from parts
    of this source image



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
* fast text rendering, with unicode
* load ttf font

* test rendering by pixel compare reference images
* dump render result to disk, using the image lib github.com/disintegration/imaging

* speed: benchmark time to render 100 frames
* speed: view tracks dirtiness of components (if they need to be redrawn)





# old ebui todo: pre 30 sep 2016:
# TODO

todo 0 :
  tests!!!
  100% coverage


[ ] todo: scene manager, se ebiten/example/blocks.go
    kunna transistera från start menu, till in-game
    
[ ] todo: om man trycker ESC in-game så ska main menu öppnas

[x] todo 1: create a main menu using MenuList
    scrolling background thingy
    
[ ] todo: resizable container with icons and interactions with icons,
    like in wayland

---


# TODO: Icon
  -> clickable button with x,y,w,h, normal and hover icon
  -> when clicked, shows a MenuList centered on screen


# Menu.List
  -> should be usable as main menu thingy
  -> ordered list of MenuItem's

# Menu.Item
  -> single item of a menu
  -> title (text, render with font)
