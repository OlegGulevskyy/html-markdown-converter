package converter

import (
	"fmt"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func tutoContainerRule() md.Rule {
	return md.Rule{
		Filter: []string{"div"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			id, exists := selec.Attr("id")
			if !exists {
				return nil
			}

			var result string
			var found bool

			if id == "tutoContainer" {
				selec.Find("div").Each(func(i int, s *goquery.Selection) {
					class, exists := s.Attr("class")
					if !exists {
						return
					}

					if class == "tutoSection" {
						s.Children().Filter("p").Each(func(i int, sel *goquery.Selection) {
							pHtml, err := sel.Html()
							if err != nil {
								return
							}
							found = true
							result += pHtml
						})
					}
				})
			}

			withWhatIsNext := whatIsNextWrapper(strings.TrimSpace(result))

			if found {
				return md.String(withWhatIsNext)
			}
			return nil
		},
	}
}

func whatIsNextWrapper(content string) string {
	return fmt.Sprintf(`

:::tip What's next

%s

:::`, content)
}
