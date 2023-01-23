package converter

import (
	"fmt"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func alertsRules() md.Rule {
	return md.Rule{
		Filter: []string{"div"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			className, exists := selec.Attr("class")
			if !exists {
				return nil
			}

			selec.Children().Each(func(i int, s *goquery.Selection) {
				if s.Is("em") && s.HasClass("material-icons") {
					s.ReplaceWith("")
				}
			})

			// escape first element if contains {{ as it messes up with JSX rendering
			firstEl := selec.First()
			firstHtml, err := firstEl.Html()
			if err != nil {
				return nil
			}
			r, _ := regexp.Compile("{{(.*)}}")
			markerRegexp := r.FindStringSubmatch(firstHtml)

			if len(markerRegexp) > 0 {
				marker := markerRegexp[0]
				escapedMarker := fmt.Sprintf("{`%s`}", marker)
				updatedHtmlFirstElement := strings.Replace(firstHtml, marker, escapedMarker, 1)
				firstEl.SetHtml(updatedHtmlFirstElement)
			}

			updatedHtml, err := selec.Html()
			updatedHtml = strings.TrimSpace(updatedHtml)
			if err != nil {
				return nil
			}

			if className == "alert alert-info" {
				element := fmt.Sprintf(":::note\n%s\n:::", updatedHtml)
				return &element
			}

			if className == "alert alert-warning" {
				element := fmt.Sprintf(":::caution Warning\n%s\n:::", updatedHtml)
				return &element
			}

			if className == "alert alert-danger" {
				element := fmt.Sprintf(":::danger\n%s\n:::", updatedHtml)
				return &element
			}

			if className == "alert alert-success" {
				element := updatedHtml
				return &element
			}
			return nil
		},
	}
}
