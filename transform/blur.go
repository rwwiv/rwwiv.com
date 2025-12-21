package transform

import (
	"errors"
	"image"

	"github.com/esimov/stackblur-go"
)

// Blur applies a blur effect to an image.
type Blur struct {
	Strength uint32
}

// NewBlur creates a new Blur transform from parameters.
func NewBlur(params map[string]any) (Transform, error) {
	strength, ok := params["blur_strength"]
	if !ok {
		return nil, errors.New("blur_strength parameter required")
	}

	var s uint32
	switch v := strength.(type) {
	case float64:
		s = uint32(v)
	case int:
		s = uint32(v)
	default:
		return nil, errors.New("blur_strength must be a number")
	}

	return &Blur{Strength: s}, nil
}

// Apply applies the blur effect to the image.
func (b *Blur) Apply(img image.Image) (image.Image, error) {
	return stackblur.Process(img, b.Strength)
}
