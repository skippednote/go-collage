package handler

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"

	"github.com/skippednote/collage/download"
	"github.com/skippednote/collage/drawimage"
	"github.com/skippednote/collage/imagemanipulation"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	pictures, err := download.GetPictures("https://www.axelerant.com/about", `<div class="emp-avatar">\s+<img src="(.+jpg)\?.+" width="300"`)
	if err != nil {
		fmt.Println("Failed to download", err)
	}
	collage := drawimage.Drawimage(pictures)
	manipulatedCollage := imagemanipulation.Manipulate(collage)
	buf := &bytes.Buffer{}
	jpeg.Encode(buf, manipulatedCollage, nil)
	pictures = nil

	w.Header().Add("cache-control", "max-age=3600, public")
	w.Write(buf.Bytes())
}
