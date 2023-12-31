package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ivanatias/imgdl/colors"
	"github.com/ivanatias/imgdl/utils"
)

func main() {
	var from, to string
	now := time.Now()

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

	err = os.MkdirAll(to, 0755)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup

	imgCount := 0

	for scanner.Scan() {
		imageUrl := scanner.Text()

		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			utils.DownloadAndSave(url, to, &imgCount)

		}(imageUrl)
	}

	wg.Wait()

	elapsed := time.Since(now).Milliseconds()

	colors.Green.Printf("\n%d images saved on %s (%dms)\n", imgCount, to, elapsed)
}
