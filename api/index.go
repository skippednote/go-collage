package handler

import (
	"bytes"
	"encoding/json"
	"image/jpeg"
	"net/http"

	"github.com/skippednote/collage/download"
	"github.com/skippednote/collage/drawimage"
	"github.com/skippednote/collage/imagemanipulation"
)

type Form struct {
	Uri   string `json:"uri"`
	Regex string `json:"regex"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	collage := drawimage.Drawimage(pictures)
	manipulatedCollage, err := imagemanipulation.Manipulate(collage, gray, width)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := &bytes.Buffer{}
	err = jpeg.Encode(buf, manipulatedCollage, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("cache-control", "max-age=3600, public")
	w.Write(buf.Bytes())
}
