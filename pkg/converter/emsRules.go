package converter

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func emsRule() md.Rule {
	return md.Rule{
		Filter: []string{"em"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			empty := ""
			return &empty
		},
	}
}
