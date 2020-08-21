package drawimage

import (
	"image"
	"image/draw"
	"sync"
	"sync/atomic"

	"github.com/skippednote/collage/download"
)

func drawChunk(x, y *int32, collage *image.RGBA, data image.Image, wg *sync.WaitGroup) {
	if *x == 23 {
		atomic.SwapInt32(x, 0)
		atomic.AddInt32(y, 1)
	} else {
		atomic.AddInt32(x, 1)
	}
	imgRect := image.Rectangle{
		Min: image.Point{
			X: int(*x * 300),
			Y: int(*y * 300),
		},
		Max: image.Point{
			X: int(*x*300 + 300),
			Y: int(*y*300 + 300),
		},
	}
	draw.Draw(collage, imgRect, data, image.Point{0, 0}, draw.Src)
	wg.Done()
}

func Drawimage(pictures []download.PictureData) *image.RGBA {
	picturesLen := len(pictures)
	extra := picturesLen % 24
	if extra > 0 {
		extra = 24 - extra
	}
	pictures = append(pictures, pictures[0:extra]...)
	picturesLen = len(pictures)
	collageRect := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 300 * 24,
			Y: int(300 * picturesLen / 24),
		},
	}
	collage := image.NewRGBA(collageRect)

	var wg sync.WaitGroup
	var x int32 = -1
	var y int32 = 0
	for _, picture := range pictures {
		wg.Add(1)
		go drawChunk(&x, &y, collage, picture.Data, &wg)
	}

	return collage
}
