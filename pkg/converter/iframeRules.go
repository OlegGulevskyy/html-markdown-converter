package converter

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func iframeRule() md.Rule {
	return md.Rule{
		Filter: []string{"div"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			fmt.Println("Iframe rule", content)
			dataType, exists := selec.Attr("data-type")
			if !exists {
				return nil
			}

			if dataType != "AwesomeTableView" {
				return nil
			}

			dataViewId, exists := selec.Attr("data-viewid")
			if !exists {
				return nil
			}

			iframeEl := fmt.Sprintf(`<iframe referrerpolicy="no-referrer-when-downgrade" data-type="AwesomeTableView" name="toto" src="https://view-awesome-table.com/%s/view/" scrolling="no"></iframe>`, dataViewId)

			return &iframeEl
		},
	}
}
