package main

import (
	"net/http"

	"github.com/pkg/profile"
	"github.com/skippednote/go-collage/route"
)

var Handler = route.Handler

func main() {
	defer profile.Start().Stop()
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
