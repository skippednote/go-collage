package main

import (
	"github.com/skippednote/collage/download"
	"github.com/skippednote/collage/imagemanipulation"
)

func main() {
	a := download.DownloadAvatars()
	c := combineAvatars(a)
	collage := imagemanipulation.Manipulate(c)
}
