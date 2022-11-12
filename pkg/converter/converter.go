package converter

import (
	"fmt"
	"log"
	"path/filepath"
	"migration-helper/pkg/files"
	"migration-helper/pkg/images"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type ConvertResult struct {
	Markdown string
	Images   []images.Image
}

func Convert(str string, articleName string, imagesDestination string) ConvertResult {
	imgs := []images.Image{}

	handleImagesRule := md.Rule{
		Filter: []string{"img"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			log.Println(content)
			src, _ := selec.Attr("src")
			alt, _ := selec.Attr("alt")
			width, _ := selec.Attr("width")
			height, _ := selec.Attr("height")

			destPath := filepath.Join(imagesDestination, articleName)
			imgName := filepath.Base(src)
			importSrc := fmt.Sprintf("./%s/%s", destPath, imgName)

			// extract images for future processing (downloading)
			finalImg := images.Image{
				Src:             src,
				Alt:             alt,
				Width:           width,
				Height:          height,
				DestinationPath: destPath,
				Name:            imgName,
				ImportPath:      importSrc,
				ImportName:      files.TransformToImportName(imgName),
			}

			imgs = append(imgs, finalImg)

			var widthValue string
			var heightValue string
			if width == "" {
				widthValue = images.WIDTH
			}

			if height == "" {
				heightValue = images.HEIGHT
			}

			text := fmt.Sprintf(
				"<img src={%s} alt=\"%s\" style={{ width: \"%spx\", height: \"%spx\" }} />",
				finalImg.ImportName,
				alt,
				widthValue,
				heightValue,
			)

			return &text
		},
	}
	conv := md.NewConverter("", true, nil)

	conv.AddRules(handleImagesRule)

	markdown, err := conv.ConvertString(str)
	if err != nil {
		log.Fatalf("Error with converter: %v", err)
	}

	return ConvertResult{
		Markdown: markdown,
		Images:   imgs,
	}
}
