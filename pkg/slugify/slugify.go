package slugify

import (
	"regexp"
	"strings"
)

func Slugify(input string) string {
	lowercase := strings.ToLower(input)

	regex := regexp.MustCompile("[^a-z0-9]+")
	cleaned := regex.ReplaceAllString(lowercase, "-")

	result := strings.Trim(cleaned, "-")

	return result
}
