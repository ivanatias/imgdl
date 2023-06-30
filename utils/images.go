package utils

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func getFilename(urlString string) (string, error) {
	parsedUrl, err := url.Parse(urlString)

	if err != nil {
		return "", err
	}

	path := parsedUrl.Path

	const pattern = `[^/]+\.[^/]+`

	reg := regexp.MustCompile(pattern)
	matches := reg.FindStringSubmatch(path)

	if len(matches) == 0 {
		return "", nil
	}

	return matches[0], nil
}

func DownloadAndSave(resourceUrl, savePath string) {
	response, err := http.Get(resourceUrl)

	if err != nil {
		Red.Println("Unable to download:", resourceUrl)

		return
	}

	defer response.Body.Close()

	err = os.MkdirAll(savePath, 0755)

	if err != nil {
		panic(err)
	}

	contentType := response.Header.Get("Content-Type")
	isImage := strings.Split(contentType, "/")[0] == "image"

	if !isImage {
		Yellow.Printf(
			"Skipping resource %s because it's not an image\n",
			resourceUrl,
		)

		return
	}

	imageData, err := io.ReadAll(response.Body)

	if err != nil {
		Red.Println("Unable to read image:", resourceUrl)

		return
	}

	imageFilename, err := getFilename(resourceUrl)

	if err != nil || len(imageFilename) == 0 {
		Red.Println("Unable to save image from:", resourceUrl)

		return
	}

	err = os.WriteFile(filepath.Join(savePath, imageFilename), imageData, 0644)

	if err != nil {
		Red.Println("Unable to save image from:", resourceUrl)

		return
	}

	Cyan.Printf("Saved %s\n", imageFilename)
}
