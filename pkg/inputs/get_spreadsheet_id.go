package inputs

import "errors"

func getSpreadsheetIdQuestion() question {
	return question{
		validate: func(input string) error {
			if input == "" {
				return errors.New("Must provide Spreadsheet ID.")
			}
			return nil
		},
		label: "Spreadsheet ID",
		saveResult: func(res string, input *UserInputs) {
			input.SpreadsheetID = res
		},
	}
}
