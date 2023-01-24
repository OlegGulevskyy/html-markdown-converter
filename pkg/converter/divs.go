package converter

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	plug "github.com/JohannesKaufmann/html-to-markdown/plugin"
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
				markdown := reconvertHtmlToMd(updatedHtml)
				element := fmt.Sprintf("\n:::note\n%s\n:::", markdown)
				return &element
			}

			if className == "alert alert-warning" {
				markdown := reconvertHtmlToMd(updatedHtml)
				element := fmt.Sprintf("\n:::caution Warning\n%s\n:::", markdown)
				return &element
			}

			if className == "alert alert-danger" {
				markdown := reconvertHtmlToMd(updatedHtml)
				element := fmt.Sprintf("\n:::danger\n%s\n:::", markdown)
				return &element
			}

			if className == "alert alert-success" {
				markdown := reconvertHtmlToMd(updatedHtml)
				return &markdown
			}
			return nil
		},
	}
}

// running converter on the HTML that has been replaced
// html-to-markdown lib will ignore it otherwise
// but, ignore any images
func reconvertHtmlToMd(htmlContent string) string {
	conv := md.NewConverter("", true, nil)
	conv.Use(plug.GitHubFlavored())

	conv.AddRules(
		alertsRules(),
		copyToCbRule(),
		escapingSoloTagsRules(),
		tutoContainerRule(),
		escapeSingleTagsRule(),
		tabsRule(),
	)

	markdown, err := conv.ConvertString(htmlContent)
	if err != nil {
		log.Fatalf("Error with re-converter: %v", err)
	}

	return markdown
}
