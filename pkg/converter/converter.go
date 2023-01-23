package converter

import (
	"log"
	"migration-helper/pkg/images"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type ConvertResult struct {
	Markdown string
	Images   []images.Image
}

func Convert(str string, categoryName string, imagesDestination string) ConvertResult {
	// collection of all images to be used in Imports in the execution later
	imgs := []images.Image{}
	handleImagesRule := imagesRule(imagesDestination, categoryName, &imgs)

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
