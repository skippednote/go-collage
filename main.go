package main

import (
	"github.com/skippednote/collage/download"
	"github.com/skippednote/collage/drawimage"
	"github.com/skippednote/collage/imagemanipulation"
)

func main() {
	// defer profile.Start().Stop()
	avatars := download.DownloadAvatars()
	collage := drawimage.Drawimage(avatars)
	imagemanipulation.Manipulate(collage)
}
