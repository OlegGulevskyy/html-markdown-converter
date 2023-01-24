package converter

import (
	"log"
	"migration-helper/pkg/images"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

type ConvertResult struct {
	Markdown string
	Images   []images.Image
}

func Convert(str string, categoryName string, articleName string, imagesDestination string) ConvertResult {
	// collection of all images to be used in Imports in the execution later
	imgs := []images.Image{}
	conv := md.NewConverter("", true, nil)
	conv.Use(plugin.GitHubFlavored())

	conv.AddRules(
		imagesRule(imagesDestination, categoryName, articleName, &imgs),
		alertsRules(),
		copyToCbRule(),
		escapingSoloTagsRules(),
		tutoContainerRule(),
		escapeSingleTagsRule(),
		tabsRule(),
	)

	markdown, err := conv.ConvertString(str)
	if err != nil {
		log.Fatalf("Error with converter: %v", err)
	}

	return ConvertResult{
		Markdown: markdown,
		Images:   imgs,
	}
}
