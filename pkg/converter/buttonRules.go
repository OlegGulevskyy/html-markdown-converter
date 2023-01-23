package converter

import (
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func copyToCbRule() md.Rule {
	return md.Rule{
		Filter: []string{"button"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			className, exists := selec.Attr("class")
			if !exists {
				return nil
			}

			if className == "copyButton" {
				selec.SetHtml("")
			}
			updatedHtml, err := selec.Html()
			updatedHtml = strings.TrimSpace(updatedHtml)
			if err != nil {
				return nil
			}

			return &updatedHtml
		},
	}
}
