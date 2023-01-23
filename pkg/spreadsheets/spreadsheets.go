package spreadsheets

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func newSpreadsheetService(creds []byte) *sheets.Service {
	// Service account based oauth2 two legged integration
	ctx := context.Background()
	srv, err := sheets.NewService(
		ctx,
		option.WithCredentialsJSON(creds),
		option.WithScopes(sheets.SpreadsheetsScope),
	)

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
		return nil
	}

	return srv
}

func GetSpreadsheetValuesBulk(ssId string, creds []byte, ranges []string) []*sheets.ValueRange {
	srv := newSpreadsheetService(creds)
	resp, err := srv.Spreadsheets.Values.BatchGet(ssId).Ranges(ranges...).Do()
	if err != nil {
		log.Fatalf("Unable to do bulk request to retrieve data from sheet: %v", err)
	}

	return resp.ValueRanges
}

func GetSpreadsheetValues(ssId string, sheetRange string, creds []byte) [][]interface{} {
	srv := newSpreadsheetService(creds)
	resp, err := srv.Spreadsheets.Values.Get(ssId, sheetRange).Do()
	if err != nil {
		log.Fatalf("error getting spreadsheet: %v", err)
	}

	return resp.Values
}

func GetSpreadsheetTitle(ssId string, creds []byte) (string, error) {
	srv := newSpreadsheetService((creds))

	resp, err := srv.Spreadsheets.Get(ssId).Do()
	if err != nil {
		return "", err
	}

	return resp.Properties.Title, nil
}
