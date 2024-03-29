package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"migration-helper/pkg/converter"
	"migration-helper/pkg/files"
	"migration-helper/pkg/images"
	"migration-helper/pkg/spreadsheets"
)

const (
	htmlIndex           = 0
	actualRangeIndex    = 1
	colNameIndexInRange = 0
)

func getTestFile() string {
	f, err := os.ReadFile("./test.html")
	if err != nil {
		log.Panicln(err)
	}

	return string(f)
}

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func getServiceAccountCreds() []byte {
	c, err := creds.ReadFile("credentials.json")
	if err != nil {
		log.Fatalln("Cannot use imported credentials file")
	}

	return c
}

type OperationRunStatus struct {
	IsDone                         bool    `json:"is_done"`
	ImagesRequestedForProcessing   int     `json:"images_requested_for_processing"`
	ArticlesRequestedForProcessing int     `json:"articles_requested_for_processing"`
	TimeElapsed                    float64 `json:"time_elapsed"`
}

type RunProps struct {
	SpreadsheetId          string `json:"spreadsheet_id"`
	HtmlBodyRange          string `json:"html_body_range"`
	ArticleNameSheetRange  string `json:"article_name_sheet_range"`
	CategoryNameSheetRange string `json:"category_name_sheet_range"`
	MdDestFolder           string `json:"md_dest_folder"`
	ImagesDestFolder       string `json:"images_dest_folder"`
}

func (a *App) Run(rp RunProps) OperationRunStatus {
	start := time.Now()

	creds := getServiceAccountCreds()

	bulkRanges := spreadsheets.GetSpreadsheetValuesBulk(rp.SpreadsheetId, creds, []string{rp.HtmlBodyRange, rp.ArticleNameSheetRange, rp.CategoryNameSheetRange})

	rangeValuesByRangeName := func(rangeName string) ([][]interface{}, error) {
		for _, rangeValue := range bulkRanges {
			actualRange := strings.Split(rangeValue.Range, "!")[actualRangeIndex] // index 0 is the sheet name
			colName := strings.Split(actualRange, ":")[colNameIndexInRange]
			reqRangeName := strings.Split(rangeName, "!")[actualRangeIndex]
			if strings.Contains(reqRangeName, colName) {
				return rangeValue.Values, nil
			}
		}
		return nil, errors.New("range not found")
	}

	htmlBodiesRows, err := rangeValuesByRangeName(rp.HtmlBodyRange)
	checkError(err)
	titlesRows, err := rangeValuesByRangeName(rp.ArticleNameSheetRange)
	checkError(err)
	categoryRows, err := rangeValuesByRangeName(rp.CategoryNameSheetRange)
	checkError(err)

	if htmlBodiesRows == nil {
		log.Fatalf("do not have any rows fetched")
	}

	var allImgs []images.Image

	for i, row := range htmlBodiesRows {
		html := row[htmlIndex]
		name := titlesRows[i][htmlIndex]
		categoryName := categoryRows[i][htmlIndex]

		sanitizedName := files.SanitizeFileName(fmt.Sprint(name))
		importsPerArticle := []string{}

		res := converter.Convert(
			fmt.Sprint(html),
			fmt.Sprint(categoryName),
			fmt.Sprint(name),
			rp.ImagesDestFolder,
			&importsPerArticle,
		)

		newpath := articlePath(rp.MdDestFolder, fmt.Sprint(categoryName))
		err := os.MkdirAll(newpath, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to create folder: %v", err)
		}

		fileName := fmt.Sprintf("%s/%s.mdx", newpath, sanitizedName)

		if len(res.Images) > 0 {
			importsPerArticle = append(importsPerArticle, "import useBaseUrl from '@docusaurus/useBaseUrl';")
		}

		// construct the article .mdx
		var finalMd string

		// first - prepend all the article imports
		for _, imp := range importsPerArticle {
			finalMd = finalMd + fmt.Sprintf("%s\n", imp)
		}

		// add frontmatter such as
		// ---
		// title: "Name of the article"
		// ---
		finalMd = addFrontMatter(finalMd, fmt.Sprint(name))

		// add the actual article content in md format
		finalMd = finalMd + "\n\n" + res.Markdown
		files.SaveStringAsFile(fileName, finalMd)

		log.Println("SAVED FILE", fileName)
		fmt.Println("Retrieve images for article", len(res.Images))

		allImgs = append(allImgs, res.Images...)
	}
	fetchedImages := images.FetchAll(allImgs)

	return OperationRunStatus{
		IsDone:                         true,
		ImagesRequestedForProcessing:   fetchedImages,
		ArticlesRequestedForProcessing: len(htmlBodiesRows),
		TimeElapsed:                    time.Since(start).Seconds(),
	}
}

func (a *App) FetchSpreadsheetInfo(ssId string) (string, error) {
	creds := getServiceAccountCreds()
	ssTitle, err := spreadsheets.GetSpreadsheetTitle(ssId, creds)
	if err != nil {
		return "", err
	}

	return ssTitle, nil
}

func articlePath(basePath string, category string) string {
	return filepath.Join(basePath, category)
}

func addFrontMatter(md string, title string) string {
	return fmt.Sprintf(`---
title: "%s"
---
`, title) + md
}

func addUseBaseUrlImport(md string) string {
	return md + "\nimport useBaseUrl from '@docusaurus/useBaseUrl';"
}
