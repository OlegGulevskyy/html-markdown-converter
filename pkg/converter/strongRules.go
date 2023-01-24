package converter

import (
	"fmt"
	"regexp"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func escapingSoloTagsRules() md.Rule {
	return md.Rule{
		// at the moment I only know that <strong> tags might contain single tags
		// example: <strong><style></strong>
		Filter: []string{"strong"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {

			// escape all single tags
			r, _ := regexp.Compile("<[a-z]+>")
			if r.MatchString(content) {
				el := fmt.Sprintf("`%s`", content)
				return &el
			}

			// delete Procedure and Result strong tags
			if content == "Procedure" || content == "Result" {
				selec.SetHtml("")

				updatedHtml, err := selec.Html()
				if err != nil {
					return nil
				}
				return &updatedHtml
			}

			return nil
		},
	}
}
