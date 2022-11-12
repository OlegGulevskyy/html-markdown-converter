package images

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
)

type Image struct {
	Src             string
	Alt             string
	Width           string
	Height          string
	DestinationPath string
	Name            string
	ImportPath      string
	ImportName      string
}

const (
	WIDTH  = "200"
	HEIGHT = "200"
)

func (i *Image) FetchFromUrl() {

	res, err := http.Get(i.Src)
	if err != nil {
		message := fmt.Sprintf("Could not fetch image via URL - %s", i.Src)
		log.Println("error", message, err)
	}
	defer res.Body.Close()

	newpath := filepath.Join(".", i.DestinationPath)
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Println("error making dirs", err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", newpath, path.Base(i.Src)))

	if err != nil {
		log.Println("error", "cold not save file to the drive")
		log.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)

	if err != nil {
		log.Println("error", "failed to copy image data to file")
	}

	fmt.Println("Success!")
}

func FetchAll(images []Image) {
	wg := sync.WaitGroup{}
	for _, image := range images {
		log.Println("Fething images by URL...", image.Src)
		wg.Add(1)
		go func(img Image) {
			img.FetchFromUrl()
			wg.Done()
		}(image)
	}
	wg.Wait()
}
