package main

import (
	"net/http"

	"github.com/skippednote/go-collage/api"
)

func main() {
	// defer profile.Start().Stop()
	http.HandleFunc("/", api.Handler)
	http.ListenAndServe(":8080", nil)
}
