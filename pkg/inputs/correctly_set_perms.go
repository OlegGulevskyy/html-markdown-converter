package inputs

func getHaveCorrectlySetPermissionsQuestion() question {
	return question{
		isConfirm: true,
		label:     "Have you shared the Spreadsheet with me? ('html-to-markdown-extractor@dev-form-publisher.iam.gserviceaccount.com')",
		saveResult: func(res string, input *UserInputs) {
			input.HaveCorrectlySetPermissions = res
		},
	}
}
