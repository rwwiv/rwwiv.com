package transform

import (
	"errors"
	"image"
	"image/color"
	"math"
)

// Pixelate applies a pixelation effect to an image.
type Pixelate struct {
	PixelSize int
}

// NewPixelate creates a new Pixelate transform from parameters.
func NewPixelate(params map[string]any) (Transform, error) {
	size, ok := params["pixel_size"]
	if !ok {
		return nil, errors.New("pixel_size parameter required")
	}

	var s int
	switch v := size.(type) {
	case float64:
		s = int(v)
	case int:
		s = v
	default:
		return nil, errors.New("pixel_size must be a number")
	}

	return &Pixelate{PixelSize: s}, nil
}

// Apply applies the pixelation effect to the image.
func (p *Pixelate) Apply(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	pixelSize := p.PixelSize

	pixelatedImg := image.NewRGBA(bounds)

	// if pixelSize is 1, just copy the image
	if pixelSize <= 1 {
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				pixel := img.At(x, y)
				col := color.RGBAModel.Convert(pixel).(color.RGBA)
				pixelatedImg.Set(x, y, col)
			}
		}
		return pixelatedImg, nil
	}

	var sizeFillsImg bool
	// clamp pixelSize to image size
	if width > height && pixelSize > width {
		pixelSize = width
		sizeFillsImg = true
	} else if height >= width && pixelSize > height {
		pixelSize = height
		sizeFillsImg = true
	}

	// if pixelSize is as large as the image, just fill the image with the average color
	if sizeFillsImg {
		r, g, b := meanAvgColor(img)
		col := color.RGBA{r, g, b, 255}
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				pixelatedImg.Set(x, y, col)
			}
		}
		return pixelatedImg, nil
	}

	centerX := width / 2
	centerY := height / 2

	// start at the top left corner of the image
	startX := centerX % pixelSize
	if startX != 0 {
		startX -= pixelSize
	}
	startY := centerY % pixelSize
	if startY != 0 {
		startY -= pixelSize
	}

	for x := startX; x < width; x += pixelSize {
		for y := startY; y < height; y += pixelSize {
			rect := image.Rect(x, y, x+pixelSize, y+pixelSize)

			// clamp rect to image bounds
			if rect.Min.X < 0 {
				rect.Min.X = 0
			}
			if rect.Min.Y < 0 {
				rect.Min.Y = 0
			}
			if rect.Max.X > width {
				rect.Max.X = width
			}
			if rect.Max.Y > height {
				rect.Max.Y = height
			}

			r, g, b := meanAvgColorOverRect(img, rect)
			col := color.RGBA{r, g, b, 255}

			// fill rect with average color
			for x2 := rect.Min.X; x2 <= rect.Max.X; x2++ {
				for y2 := rect.Min.Y; y2 <= rect.Max.Y; y2++ {
					pixelatedImg.Set(x2, y2, col)
				}
			}
		}
	}

	return pixelatedImg, nil
}

func meanAvgColor(img image.Image) (red, green, blue uint8) {
	bounds := img.Bounds()
	return meanAvgColorOverRect(img, bounds)
}

func meanAvgColorOverRect(img image.Image, rect image.Rectangle) (red, green, blue uint8) {
	var redSum float64
	var greenSum float64
	var blueSum float64

	// sum up all the colors in the rect
	for x := rect.Min.X; x <= rect.Max.X; x++ {
		for y := rect.Min.Y; y <= rect.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			redSum += float64(r/256) * float64(r/256)
			greenSum += float64(g/256) * float64(g/256)
			blueSum += float64(b/256) * float64(b/256)
		}
	}

	rectArea := float64((rect.Dx() + 1) * (rect.Dy() + 1))

	// average the sums
	red = uint8(math.Round(math.Sqrt(redSum / rectArea)))
	green = uint8(math.Round(math.Sqrt(greenSum / rectArea)))
	blue = uint8(math.Round(math.Sqrt(blueSum / rectArea)))

	return
}
