package imagemanipulation

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func Manipulate() {
	f, err := os.Open("./avatars/arun.jpg")
	if err != nil {
		log.Fatalln("Failed to read the file")
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatalln("Failed to decode the file")
	}
	bounds := img.Bounds()

	grayImage := image.NewGray(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			oldPixel := img.At(x, y)
			newPixel := color.GrayModel.Convert(oldPixel)
			grayImage.Set(x, y, newPixel)
		}
	}
	resizedImage := resize.Thumbnail(100, 100, grayImage, resize.Lanczos3)

	f, err = os.Create("resized-gray-arun.jpg")
	if err != nil {
		log.Fatalln("Failed to create the gray image")
	}
	defer f.Close()
	err = jpeg.Encode(f, resizedImage, nil)
	if err != nil {
		log.Fatalln("Failed to encode the gray image")
	}
}
