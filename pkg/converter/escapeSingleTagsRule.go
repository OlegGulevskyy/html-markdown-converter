package converter

import (
	"fmt"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

var escapeTagContainers = []string{"p"}
var escapeTags = []string{"script", "style", "head", "div"}

func escapeSingleTagsRule() md.Rule {
	return md.Rule{
		Filter: escapeTagContainers,
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			var foundAnyTag bool

			for _, tag := range escapeTags {
				r, _ := regexp.Compile(fmt.Sprintf("<(%s)>", tag))
				tagRegexp := r.FindStringSubmatch(content)

				if len(tagRegexp) > 0 {
					foundAnyTag = true
					tag := tagRegexp[0]
					escapedTag := fmt.Sprintf("`%s`", tag)

					content = strings.Replace(content, tag, escapedTag, 1)
				}
			}
			if foundAnyTag {
				return md.String(content)
			}
			return nil
		},
	}
}
