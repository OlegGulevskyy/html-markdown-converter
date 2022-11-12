package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"migration-helper/pkg/converter"
	"migration-helper/pkg/files"
	"migration-helper/pkg/images"
	"migration-helper/pkg/spreadsheets"
)

const (
	htmlIndex = 0
)

func getTestFile() string {
	f, err := os.ReadFile("./test.html")
	if err != nil {
		log.Panicln(err)
	}

	return string(f)
}

func getServiceAccountCreds() []byte {
	c, err := creds.ReadFile("credentials.json")
	if err != nil {
		log.Fatalln("Cannot use imported credentials file")
	}

	return c
}

func (a *App) Run(spreadsheetId string, htmlBodyRange string, articleNameSheetRange string) {

	creds := getServiceAccountCreds()

	htmlBodiesRows := spreadsheets.GetSpreadsheetValues(spreadsheetId, htmlBodyRange, creds)
	titlesRows := spreadsheets.GetSpreadsheetValues(spreadsheetId, articleNameSheetRange, creds)

	if htmlBodiesRows == nil {
		log.Fatalf("do not have any rows fetched")
	}

	var allImgs []images.Image

	for i, row := range htmlBodiesRows {
		html := row[htmlIndex]
		name := titlesRows[i][htmlIndex]

		sanitizedName := files.SanitizeFileName(fmt.Sprint(name))

		res := converter.Convert(fmt.Sprint(html), sanitizedName)

		newpath := filepath.Join(".", "output")
		err := os.MkdirAll(newpath, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to create folder: %v", err)
		}

		fileName := fmt.Sprintf("%s/%s.mdx", newpath, sanitizedName)
		var finalMd string
		for _, img := range res.Images {
			finalMd = finalMd + fmt.Sprintf("import %s from \"%s\"\n", img.ImportName, img.ImportPath)
		}
		finalMd = finalMd + "\n\n" + res.Markdown

		files.SaveStringAsFile(fileName, finalMd)

		log.Println("SAVED FILE", fileName)

		allImgs = append(allImgs, res.Images...)

	}

	images.FetchAll(allImgs)

	log.Println("Done, merci, grous bisous!")

}

func (a *App) FetchSpreadsheetInfo(ssId string) (string, error) {
	creds := getServiceAccountCreds()
	ssTitle, err := spreadsheets.GetSpreadsheetTitle(ssId, creds)
	if err != nil {
		return "", err
	}

	return ssTitle, nil

}
