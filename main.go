package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func getHTML() []byte {
	res, err := http.Get("https://www.axelerant.com/about")
	if err != nil {
		log.Fatalln("Failed to get the about page.", err.Error())
	}

	defer res.Body.Close()
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Failed to read the response body.", err.Error())
	}

	return html
}

func getAvatars(html []byte) []string {
	var re = regexp.MustCompile(`<div class="emp-avatar">\s+<img src="(.+jpg)\?.+" width="300"`)

	reResult := re.FindAllSubmatch(html, 200)
	avatars := make([]string, len(reResult))

	for _, v := range reResult {
		avatars = append(avatars, string(v[1]))
	}

	return avatars
}

func downloadsDirectory() {
	err := os.RemoveAll("avatars")
	if err != nil {
		log.Fatalln("Failed to delete the avatars directory")
	}
	err = os.Mkdir("avatars", 0766)
	if err != nil {
		log.Fatalln("Failed to create the avatars directory")
	}
}

func downloadImage(image string, wg *sync.WaitGroup) {
	defer wg.Done()
	s := strings.Split(image, "/")
	filename := strings.ToLower(s[len(s)-1])

	baseURL := "https://www.axelerant.com"
	res, err := http.Get(baseURL + image)
	if err != nil {
		log.Println("Failed to fetch the image")
	}

	defer res.Body.Close()
	imageByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Failed to read the response body")
	}

	log.Printf("Downloading the file: %s", filename)
	path := filepath.Join(".", "avatars", filename)
	err = ioutil.WriteFile(path, imageByte, 0766)
	if err != nil {
		log.Println("Failed to write the image to disk", err.Error(), image)
	}
}

func downloadImages(avatars []string) {
	var wg sync.WaitGroup
	for _, image := range avatars {
		if len(image) > 0 {
			wg.Add(1)
			go downloadImage(image, &wg)
		}
	}
	wg.Wait()
}

func main() {
	html := getHTML()
	avatars := getAvatars(html)
	downloadsDirectory()
	downloadImages(avatars)
	fmt.Println("DONE")
}
