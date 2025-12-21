package transform

import "image"

// Transform is the interface that all image transformations must implement.
type Transform interface {
	Apply(img image.Image) (image.Image, error)
}

// Registry maps transform names to constructor functions.
var Registry = map[string]func(params map[string]any) (Transform, error){
	"blur":      NewBlur,
	"grayscale": NewGrayscale,
	"invert":    NewInvert,
	"pixelate":  NewPixelate,
}
