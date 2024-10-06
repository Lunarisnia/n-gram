package strgram

import (
	"errors"
	"fmt"
	"strings"
)

// BIGRAM EXAMPLE
// RESULT
// _R, RE, ES, SU, UL, LT, T_

// FORMALITY LEVEL
// _F, FO, OR, RM, MA, AL, LI, IT, TY, Y<SP>, <SP>L, LE, EV, VE, EL, L_

// TEXT
// 0123
// n = 2
// Resulting Bigram:
// _T, TE, EX, XT, T_
// 00, 01, 12, 23, 3_

func Encode(text string, n int) ([][]string, error) {
	textLength := len(text)
	if n > textLength {
		return nil, errors.New("n cannot be larger than the text length: n < len(text)")
	}

	ngramResult := make([][]string, 0)

	if n == 1 {
		splitStr := strings.Split(text, "")
		for _, str := range splitStr {
			ngramResult = append(ngramResult, []string{str})
		}
		return ngramResult, nil
	}

	stringBuf := make([]string, 0)
	for range n {
		stringBuf = append(stringBuf, "_")
	}
	var insertBuffer func()
	i := 0
	insertBuffer = func() {
		if len(stringBuf) == n || i > textLength {
			stringBuf = stringBuf[1:]
		}
		if i < textLength {
			stringBuf = append(stringBuf, string(text[i]))
		}
		i++
	}

	for j := range textLength + 1 {
		if n > 2 && j == 0 {
			for range n - 2 {
				insertBuffer()
			}
		}
		insertBuffer()
		if len(stringBuf) < n {
			stringBuf = append(stringBuf, "_")
		}
		ngramResult = append(ngramResult, stringBuf)
	}

	return ngramResult, nil
}

func dumpBuffer(stringBuf []rune) {
	for _, r := range stringBuf {
		fmt.Printf(" %s ", string(r))
	}
	fmt.Println()
}
