package files

import (
	"path/filepath"
	"regexp"
	"strings"
)

func SanitizeFileName(s string) string {
	// replace special characters
	str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
	str = strings.ReplaceAll(str, " ", "-")
	str = strings.ToLower(str)

	return str
}

func TransformToImportName(s string) string {
	// remove extension
	name := strings.TrimSuffix(s, filepath.Ext(s))
	str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, " ")
	str = strings.Title(str)

	return strings.ReplaceAll(str, " ", "")
}
