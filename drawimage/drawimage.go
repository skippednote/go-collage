package drawimage

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func Drawimage(avatars []string) *image.RGBA {
	totalAvatars := len(avatars)
	extra := totalAvatars % 24
	if extra > 0 {
		extra = 24 - extra
	}
	avatars = append(avatars, avatars[0:extra]...)
	totalAvatars = len(avatars)
	fmt.Println(totalAvatars)
	x := 0
	y := 0
	collageRect := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 300 * 24,
			Y: int(300 * totalAvatars / 24),
		},
	}
	collage := image.NewRGBA(collageRect)

	for _, i := range avatars {
		f, err := os.Open(i)
		if err != nil {
			log.Fatalln("Failed to read the file")
		}
		defer f.Close()

		img, err := jpeg.Decode(f)
		if err != nil {
			log.Fatalln("Failed to decode the file")
		}

		imgRect := image.Rectangle{
			Min: image.Point{
				X: x * 300,
				Y: y * 300,
			},
			Max: image.Point{
				X: x*300 + 300,
				Y: y*300 + 300,
			},
		}

		draw.Draw(collage, imgRect, img, image.Point{0, 0}, draw.Src)

		if x == 23 {
			x = 0
			y += 1
		} else {
			x += 1
		}
	}

	// ouput, err := os.Create("./draw.jpg")
	// if err != nil {
	// 	log.Fatalln("Failed to write the draw image")
	// }
	// jpeg.Encode(ouput, collage, nil)
	return collage
}
