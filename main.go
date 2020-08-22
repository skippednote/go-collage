package main

import (
	"bytes"
	"encoding/json"
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

type Form struct {
	Uri   string `json:"uri"`
	Regex string `json:"regex"`
}

func main() {
	defer profile.Start().Stop()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST request is supported", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		form := &Form{}
		err := json.NewDecoder(r.Body).Decode(form)
		if err != nil {
			http.Error(w, "Failed to decode the request body", http.StatusInternalServerError)
			return
		}

		query := r.URL.Query()
		gray := query.Get("gray")
		width := query.Get("width")
		// pictures, err := download.GetPictures("https://www.axelerant.com/about", `<div class="emp-avatar">\s+<img src="(.+jpg)\?.+" width="300"`)
		pictures, err := download.GetPictures(form.Uri, form.Regex)
		if err != nil {
			fmt.Println("Failed to download", err)
		}
		collage := drawimage.Drawimage(pictures)
		manipulatedCollage, err := imagemanipulation.Manipulate(collage, gray, width)
		buf := &bytes.Buffer{}
		jpeg.Encode(buf, manipulatedCollage, nil)
		w.Header().Set("Accept", "image/jpeg")
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
