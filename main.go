package main

import (
	"bufio"
	"flag"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var (
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	cyan   = color.New(color.FgCyan)
)

func main() {
	var from string
	var to string

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	defaultTo := filepath.Join(cwd, "imgdl")

	flag.StringVar(&from, "from", "", "Path to text file with all images urls")
	flag.StringVar(
		&to,
		"to",
		defaultTo,
		"Path to folder where images will be saved",
	)
	flag.Parse()

	if len(from) == 0 {
		panic("Path to text file with all images urls is required")
	}

	fromSlice := strings.Split(from, "/")
	textFile := fromSlice[len(fromSlice)-1]
	ext := strings.Split(textFile, ".")[1]

	if ext != "txt" {
		panic("Path of file with all images urls must have .txt extension")
	}

	file, err := os.Open(from)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		imageUrl := scanner.Text()
		response, err := http.Get(imageUrl)

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		contentType := response.Header.Get("Content-Type")
		contentSlice := strings.Split(contentType, "/")

		if contentSlice[0] != "image" {
			yellow.Printf(
				"Skipping resource %s because it's not an image\n",
				imageUrl,
			)

			continue
		}

		imageData, err := io.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		imageFilename := filepath.Base(imageUrl)

		err = os.MkdirAll(to, 0755)

		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filepath.Join(to, imageFilename), imageData, 0644)

		if err != nil {
			panic(err)
		}

		cyan.Printf("Saved %s\n", imageFilename)
	}

	green.Printf("\nAll images saved on %s\n", to)
}
