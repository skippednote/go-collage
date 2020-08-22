package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/pkg/profile"
	"github.com/skippednote/collage/download"
	"github.com/skippednote/collage/drawimage"
	"github.com/skippednote/collage/imagemanipulation"
)

func main() {
	defer profile.Start().Stop()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		gray := query.Get("gray")
		width := query.Get("width")
		pictures, err := download.GetPictures("https://www.axelerant.com/about", `<div class="emp-avatar">\s+<img src="(.+jpg)\?.+" width="300"`)
		if err != nil {
			fmt.Println("Failed to download", err)
		}
		collage := drawimage.Drawimage(pictures)
		manipulatedCollage, err := imagemanipulation.Manipulate(collage, gray, width)
		buf := &bytes.Buffer{}
		jpeg.Encode(buf, manipulatedCollage, nil)
		w.Write(buf.Bytes())
	})
	http.ListenAndServe(":8080", nil)
}

func saveCollage(collage image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Failed to create the gray collage. %s", err.Error())
	}
	defer f.Close()

	err = jpeg.Encode(f, collage, nil)
	if err != nil {
		log.Printf("Failed to encode the gray collage. %s", err.Error())
	}
}
