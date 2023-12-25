package filters

import (
	"image"
	"image/color"
)

func Ordered(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	out := image.NewRGBA(bounds)

	matrix := [][]float64{
		{0, 8, 2, 10},
		{12, 4, 14, 6},
		{3, 11, 1, 9},
		{15, 7, 13, 5},
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			gray := 0.299*float64(r) + 0.587*float64(g) + 0.144*float64(b)
			index := int(gray / 256 * 6)
			if gray > matrix[x%4][y%4]*256/16 {
				index++
			}

			newColor := color.Gray{uint8(index * 16)}
			out.Set(x, y, newColor)
		}
	}
	return out.SubImage(bounds)
}
