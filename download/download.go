package download

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type PictureData struct {
	Path string
	Data image.Image
}

func getHTML(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to get the about page. %s", err.Error())
	}

	defer res.Body.Close()
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the response body. %s", err.Error())
	}

	return html, nil
}

func getPictureURLs(regex string, html []byte) ([]string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("Failed to compile the given regexp. %s", err.Error())
	}

	reResult := re.FindAllSubmatch(html, 200)
	var pictureURLs []string

	for _, v := range reResult {
		pictureURLs = append(pictureURLs, string(v[1]))
	}

	return pictureURLs, nil
}

func downloadImage(baseURL string, image string, pictures *[]PictureData, wg *sync.WaitGroup) {
	defer wg.Done()
	s := strings.Split(image, "/")
	filename := strings.ToLower(s[len(s)-1])
	path := filepath.Join(".", "avatars", filename)

	res, err := http.Get(baseURL + image)
	if err != nil {
		log.Printf("Failed to fetch the image. %s", err.Error())
	}
	defer res.Body.Close()

	img, err := jpeg.Decode(res.Body)
	if err != nil {
		log.Printf("Failed to decode the image. %s", err.Error())
	}

	// log.Printf("Downloading the file: %s", filename)
	*pictures = append(*pictures, PictureData{
		Path: path,
		Data: img,
	})
}

func downloadImages(baseURL string, avatars []string, pictures *[]PictureData) error {
	var wg sync.WaitGroup
	for _, image := range avatars {
		wg.Add(1)
		go downloadImage(baseURL, image, pictures, &wg)
	}
	wg.Wait()
	return nil
}

func GetPictures(uri string, regex string) ([]PictureData, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	baseURL := u.Scheme + u.Hostname()

	var pictures []PictureData
	html, err := getHTML(uri)
	if err != nil {
		return nil, err
	}

	pictureURLs, err := getPictureURLs(regex, html)
	if err != nil {
		return nil, err
	}

	err = downloadImages(baseURL, pictureURLs, &pictures)
	if err != nil {
		return nil, err
	}

	return pictures, nil
}
