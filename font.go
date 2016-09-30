package ui

/*
// LoadDebugFont ...
func (w *World) LoadDebugFont(fileName string) error {

	//  TODO check out github.com/hajimehoshi/ebiten/examples/common/font.go
	var err error
	w.DebugFont, err = ui.NewFont("_resources/font/topaz-8.ttf", 72, 10)
	return err
}
*/

/*
import (
	"image"
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	fontColor = color.RGBA{209, 0, 0, 255}
)

// Font represents a font resource
type Font struct {
	dpi    float64
	size   float64
	font   *truetype.Font
	drawer *font.Drawer
}

// NewFont prepares a new font resource for use
func NewFont(fontName string, size float64, dpi float64) (*Font, error) {

	f, err := ebitenutil.OpenFile(fontName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var fnt Font
	fnt.font, err = truetype.Parse(b)
	fnt.size = size
	fnt.dpi = dpi

	fnt.drawer = &font.Drawer{
		Src: image.NewUniform(fontColor),
		Face: truetype.NewFace(fnt.font, &truetype.Options{
			Size:    fnt.size,
			DPI:     fnt.dpi,
			Hinting: font.HintingFull,
		}),
	}

	return &fnt, err
}

// Print draws text using the font
func (fnt *Font) Print(text string) (*ebiten.Image, error) {

	if text == "" {
		panic("empty text!")
	}

	width := fnt.drawer.MeasureString(text).Ceil()
	height := int(fnt.size) + 2 // XXX not perfect
	fnt.drawer.Dst = image.NewRGBA(image.Rect(0, 0, width, height))

	dy := (fnt.size * fnt.dpi) / 72
	fnt.drawer.Dot = fixed.P(0, int(dy))
	fnt.drawer.DrawString(text)

	return ebiten.NewImageFromImage(fnt.drawer.Dst, ebiten.FilterNearest)
}
*/
