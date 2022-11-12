package inputs

import "errors"

func getHtmlsSheetRange() question {
	return question{
		validate: func(input string) error {
			if input == "" {
				return errors.New("Expecting Sheet range. Example: Sheet!A2:A - to select the whole column A")
			}
			return nil
		},
		label: "HTML bodies sheet range",
		saveResult: func(res string, input *UserInputs) {
			input.HtmlBodySheetRange = res
		},
	}
}

func getTitlesSheetRange() question {
	return question{
		validate: func(input string) error {
			if input == "" {
				return errors.New("Expecting Sheet range. Example: Sheet!A2:A - to select the whole column A")
			}
			return nil
		},
		label: "Titles sheet range",
		saveResult: func(res string, input *UserInputs) {
			input.ArticleNameSheetRange = res
		},
	}
}
