package tokenizer

import "strings"

func Tokenize(sanitizedString string) ([][]string, error) {
	rawTokens := strings.Split(sanitizedString, " ")
	paddedTokens := make([][]string, 0)
	for _, token := range rawTokens {
		// NOTE: its 5 because we need to create pentagram for all text
		if len(token) < 5 {
			frontPad := padWithEmptySpace(token, false)
			backPad := padWithEmptySpace(token, true)
			paddedTokens = append(paddedTokens, frontPad)
			paddedTokens = append(paddedTokens, backPad)
		} else {
			paddedTokens = append(paddedTokens, strings.Split(token, ""))
		}
	}
	return paddedTokens, nil
}

func padWithEmptySpace(token string, padAtBack bool) []string {
	desiredLength := 5
	padLength := desiredLength - len(token)
	splitToken := strings.Split(token, "")
	if padAtBack {
		for range padLength {
			splitToken = append(splitToken, "")
		}
		return splitToken
	} else {
		padding := make([]string, 0)
		for range padLength {
			padding = append(padding, "")
		}
		for _, r := range splitToken {
			padding = append(padding, string(r))
		}
		return padding
	}
}
