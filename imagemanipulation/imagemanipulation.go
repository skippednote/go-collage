package imagemanipulation

import (
	"image"
	"image/color"
	"strconv"

	"github.com/nfnt/resize"
)

func Manipulate(collage *image.RGBA, gray bool, width string) (image.Image, error) {
	resizedCollage := resize.Thumbnail(1920, 1080, collage, resize.Lanczos3)
	if len(width) > 0 {
		width, err := strconv.ParseUint(width, 10, 32)
		if err != nil {
			return nil, err
		}
		resizedCollage = resize.Thumbnail(uint(width), 1080, collage, resize.Lanczos3)
	}

	if gray {
		bounds := resizedCollage.Bounds()
		grayCollage := image.NewGray(bounds)

		for y := 0; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				oldPixel := resizedCollage.At(x, y)
				newPixel := color.GrayModel.Convert(oldPixel)
				grayCollage.Set(x, y, newPixel)
			}
		}
		return grayCollage, nil
	}

	return resizedCollage, nil
}
