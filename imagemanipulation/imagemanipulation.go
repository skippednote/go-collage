package imagemanipulation

import (
	"image"
	"image/color"

	"github.com/nfnt/resize"
)

func Manipulate(collage *image.RGBA) image.Image {
	resizedCollage := resize.Thumbnail(1920, 1080, collage, resize.Lanczos3)
	bounds := resizedCollage.Bounds()
	grayCollage := image.NewGray(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			oldPixel := resizedCollage.At(x, y)
			newPixel := color.GrayModel.Convert(oldPixel)
			grayCollage.Set(x, y, newPixel)
		}
	}

	return grayCollage
}
