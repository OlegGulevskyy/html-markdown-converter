package converter

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func tabsRule() md.Rule {
	return md.Rule{
		Filter: []string{"ul"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			className, exists := selec.Attr("class")
			if !exists {
				return nil
			}

			if className != "nav nav-tabs" {
				return nil
			}

			var tabs [][]string
			selec.Find("li").Each(func(i int, s *goquery.Selection) {
				link := s.Children().First()
				href, _ := link.Attr("href")
				text := link.Text()
				tabs = append(tabs, []string{href, text})
			})

			fmt.Println("TABS => ", tabs)
			return nil
		},
	}
}
