package inputs

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"
)

type UserInputs struct {
	SpreadsheetID               string
	HtmlBodySheetRange          string
	ArticleNameSheetRange       string
	HaveCorrectlySetPermissions string
}

type question struct {
	validate   func(string) error
	label      string
	saveResult func(string, *UserInputs)
	isConfirm  bool
}

var questions = []question{
	getSpreadsheetIdQuestion(),
	getHaveCorrectlySetPermissionsQuestion(),
	getHtmlsSheetRange(),
	getTitlesSheetRange(),
}

func GetUserInput(u *UserInputs) {
	for _, q := range questions {
		prompt := promptui.Prompt{
			Label:     q.label,
			Validate:  q.validate,
			IsConfirm: q.isConfirm,
		}

		res, err := prompt.Run()
		confirmed := !errors.Is(err, promptui.ErrAbort)

		if err != nil && confirmed {
			log.Fatalf("ERROR: %v\n", err)
		}

		if !confirmed {
			log.Fatalf("Well, go and shared it then, so I can access the data..")
		}

		q.saveResult(res, u)
	}
}
