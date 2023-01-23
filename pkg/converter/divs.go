package converter

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func articleHatRule() md.Rule {
	return md.Rule{
		Filter: []string{"div"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			className, exists := selec.Attr("class")
			if exists && className == "article-chapeau" {
				element := fmt.Sprintf("_%s_", content)
				return &element
			}
			return nil
		},
	}
}
