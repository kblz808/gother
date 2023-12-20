package filters

import (
	"image"
	"image/color"
)

func Floyd(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	for y := bounds.Min.Y; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			oldPixel := grayImg.GrayAt(x, y)
			newPixel := color.Gray{Y: 0}
			if oldPixel.Y > 128 {
				newPixel.Y = 255
			}
			grayImg.SetGray(x, y, newPixel)
			err := int(oldPixel.Y) - int(newPixel.Y)
			distributeError(grayImg, x, y, err)
		}
	}

	return grayImg
}

func distributeError(img *image.Gray, x, y, err int) {
	img.SetGray(x+1, y, addError(img.GrayAt(x+1, y), err*7/16))
	img.SetGray(x-1, y+1, addError(img.GrayAt(x-1, y+1), err*3/16))
	img.SetGray(x, y+1, addError(img.GrayAt(x, y+1), err*5/16))
	img.SetGray(x+1, y+1, addError(img.GrayAt(x+1, y+1), err*1/16))
}

func addError(c color.Gray, err int) color.Gray {
	newY := int(c.Y) + err
	if newY < 0 {
		newY = 0
	} else if newY > 255 {
		newY = 255
	}
	return color.Gray{Y: uint8(newY)}
}
