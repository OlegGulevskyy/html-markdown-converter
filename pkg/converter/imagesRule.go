package converter

import (
	"fmt"
	"migration-helper/pkg/images"
	"path/filepath"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func imagesRule(imagesDestination string, categoryName string, imgs *[]images.Image) md.Rule {
	return md.Rule{
		Filter: []string{"img"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			src, _ := selec.Attr("src")
			alt, _ := selec.Attr("alt")
			width, _ := selec.Attr("width")
			height, _ := selec.Attr("height")

			destPath := filepath.Join(imagesDestination, categoryName)
			imgName := filepath.Base(src)
			importSrc := fmt.Sprintf("./%s/%s/%s", filepath.Base(imagesDestination), categoryName, imgName)

			// extract images for future processing (downloading)
			finalImg := images.Image{
				Src:             src,
				Alt:             alt,
				Width:           width,
				Height:          height,
				DestinationPath: destPath,
				Name:            imgName,
				ImportPath:      importSrc,
				ImportName:      fmt.Sprintf("useBaseUrl(\"/%s/%s\")", categoryName, imgName),
			}

			*imgs = append(*imgs, finalImg)

			element := fmt.Sprintf("<img src={%s}", finalImg.ImportName)
			element = element + fmt.Sprintf(" alt=\"%s\"", alt)
			if width != "" || height != "" {
				element = element + "style={{"
				if width != "" {
					element = element + fmt.Sprintf("width: \"%spx\",", width)
				}
				if height != "" {
					element = element + fmt.Sprintf("height: \"%spx\",", height)
				}
				element = element + " }}"
			}
			element = element + " />"

			return &element
		},
	}
}
