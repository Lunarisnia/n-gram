package tokenizer

import (
	"regexp"
	"strings"
)

func Sanitize(rawString string) string {
	solver := regexp.MustCompile("[a-zA-Z']+")
	results := solver.FindAllString(rawString, -1)
	return strings.Join(results, " ")
}
