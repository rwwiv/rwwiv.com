package transform

import (
	"image"
	"image/color"
)

// Grayscale converts an image to grayscale.
type Grayscale struct{}

// NewGrayscale creates a new Grayscale transform.
func NewGrayscale(params map[string]any) (Transform, error) {
	return &Grayscale{}, nil
}

// Apply converts the image to grayscale using ITU-R BT.601 luminance.
func (g *Grayscale) Apply(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	imgSet := image.NewRGBA(bounds)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			px := color.Gray{uint8(lum / 256)}
			imgSet.Set(x, y, px)
		}
	}
	return imgSet, nil
}
