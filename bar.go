package ui

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Bar is a progress bar (UI component)
type Bar struct {
	component
	value       int // in percent (0-100)
	borderColor color.RGBA
	fillColor   color.RGBA
	fillImage   image.Image
}

var (
	barBorderColor = color.RGBA{0x50, 0x50, 0x50, 192} // gray, 75% transparent
	barFillColor   = Yellow
)

// NewBar creates a new Bar
func NewBar(width, height int) *Bar {
	bar := Bar{}
	bar.borderColor = barBorderColor
	bar.fillColor = barFillColor
	bar.Dimension = Dimension{Width: width, Height: height}
	return &bar
}

// SetBorderColor sets the border color
func (bar *Bar) SetBorderColor(c color.RGBA) {
	bar.borderColor = c
}

// SetFillColor sets the fill color
func (bar *Bar) SetFillColor(c color.RGBA) {
	bar.fillColor = c
}

// SetFillImage sets the fill image. If set, is used instead of fill color
func (bar *Bar) SetFillImage(img image.Image) {
	bar.fillImage = img
}

// SetValue accepts a value between 0-100
func (bar *Bar) SetValue(v int) {
	if v > 100 {
		fmt.Println("warning: bar value is too high", v)
		v = 100
	}
	if v < 0 {
		fmt.Println("warning: bar value is too low", v)
		v = 0
	}
	if bar.value == v {
		return
	}
	bar.isClean = false
	bar.value = v
}

// IncValue increases value by `i` up to 100
func (bar *Bar) IncValue(i int) int {
	val := bar.value + i
	if val > 100 {
		val = 100
	}
	if val == bar.value {
		return bar.value
	}
	bar.isClean = false
	bar.value = val
	return bar.value
}

// GetValue returns the value (0-100, percent)
func (bar *Bar) GetValue() int {
	return bar.value
}

// Draw redraws internal buffer
func (bar *Bar) Draw(mx, my int) *image.RGBA {
	if bar.isHidden {
		return nil
	}
	if bar.isClean {
		return bar.Image
	}

	rect := image.Rect(0, 0, bar.Dimension.Width, bar.Dimension.Height)
	if bar.Image == nil {
		bar.Image = image.NewRGBA(rect)
	} else {
		draw.Draw(bar.Image, rect, &image.Uniform{color.Transparent}, image.ZP, draw.Src)
	}

	// draw outline
	outlineRect := image.Rect(0, 0, bar.Dimension.Width-1, bar.Dimension.Height-1)
	DrawRect(bar.Image, outlineRect, bar.borderColor)

	// convert bar.value (percent) into number of pixels to cover (width)
	pixelWidth := int(math.Floor(((float64(bar.value)/100)*float64(bar.Dimension.Width))+0.5)) - 1

	if bar.fillImage == nil {
		fillRect := image.Rect(1, 1, pixelWidth+1, bar.Dimension.Height-1)
		draw.Draw(bar.Image, fillRect, &image.Uniform{bar.fillColor}, image.ZP, draw.Src)
	} else {
		// fill using repeating image
		b := bar.fillImage.Bounds()

		for x := 0; x <= pixelWidth; x += b.Max.X {
			width := b.Max.X
			// on last image, use partial width to be pixel exact
			if x+b.Max.X >= pixelWidth {
				width = pixelWidth - x
			}
			fillRect := image.Rect(x+1, 1, x+width+1, b.Max.Y+1)
			draw.Draw(bar.Image, fillRect, bar.fillImage, image.ZP, draw.Src)
		}
	}

	bar.isClean = true
	return bar.Image
}
