package main

import (
	"net/http"

	"github.com/pkg/profile"
	"github.com/skippednote/collage/route"
)

func main() {
	defer profile.Start().Stop()
	http.HandleFunc("/", route.Handler)
	http.ListenAndServe(":8080", nil)
}
