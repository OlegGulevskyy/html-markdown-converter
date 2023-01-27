package converter

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func tabsRule(articleImports *[]string) md.Rule {
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

			if len(tabs) > 0 {
				tabImport := "import Tabs from '@theme/Tabs';"
				tabItemImport := "import TabItem from '@theme/TabItem';"

				if !contains(*articleImports, tabImport) {
					*articleImports = append(*articleImports, tabImport)
				}

				if !contains(*articleImports, tabItemImport) {
					*articleImports = append(*articleImports, tabItemImport)
				}
			}

			nextToNavTabs := selec.Next().Children()

			result := ""
			result += fmt.Sprintf("<Tabs>\n")

			for tabIndex, tab := range tabs {
				result += fmt.Sprintf("<TabItem value=\"%s\" label=\"%s\"", tab[0], tab[1])
				if tabIndex == 0 {
					result += fmt.Sprintf("default>\n")
				} else {
					result += fmt.Sprintf(">\n")
				}

				tabValue := nextToNavTabs.Eq(tabIndex)
				tabHtml, _ := tabValue.Html()

				result += fmt.Sprintf("%s\n", tabHtml)
				result += fmt.Sprintf("</TabItem>\n")
			}
			result += fmt.Sprintf("</Tabs>\n")

			return &result
		},
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
