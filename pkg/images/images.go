package images

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"

	"migration-helper/pkg/utils"
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

func (i *Image) fetchFromUrl() {

	res, err := http.Get(i.Src)
	if err != nil {
		message := fmt.Sprintf("Could not fetch image via URL - %s", i.Src)
		log.Println("error", message, err)
	}
	defer res.Body.Close()

	err = os.MkdirAll(i.DestinationPath, os.ModePerm)
	if err != nil {
		log.Println("error making dirs", err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", i.DestinationPath, path.Base(i.Src)))

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
	imagesByBatches := utils.ChunkBy(images, 150)
	log.Println(fmt.Sprintf("Broke down image array into %v", len(imagesByBatches)))

	for i, imageBatch := range imagesByBatches {
		log.Println("Image batch -> ", i)
		wg := sync.WaitGroup{}
		for _, image := range imageBatch {

			if image.Src == "" {
				fmt.Println("Image src is empty", image)
				continue
			}

			log.Println("Fething images by URL...", image.Src)
			wg.Add(1)
			go func(img Image) {
				img.fetchFromUrl()
				wg.Done()
			}(image)
		}
		wg.Wait()
	}

}
