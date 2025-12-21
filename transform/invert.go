package transform

import (
	"image"
	"image/color"
)

// Invert inverts the colors of an image.
type Invert struct{}

// NewInvert creates a new Invert transform.
func NewInvert(params map[string]any) (Transform, error) {
	return &Invert{}, nil
}

// Apply inverts all color channels of the image.
func (i *Invert) Apply(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	imgSet := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			col := color.RGBA{255 - uint8(r/256), 255 - uint8(g/256), 255 - uint8(b/256), uint8(a / 256)}
			imgSet.Set(x, y, col)
		}
	}
	return imgSet, nil
}
