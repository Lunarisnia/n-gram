package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenize(t *testing.T) {
	t.Run("Expect properly tokenized", func(t *testing.T) {
		rawToken := "hey tHere what is going ' on here wtf"
		tokens, err := Tokenize(rawToken)
		assert.Nil(t, err)
		assert.Equal(t, [][]string{
			{"h", "e", "y", "", ""},
			{"t", "H", "e", "r", "e"},
			{"w", "h", "a", "t", ""},
			{"i", "s", "", "", ""},
			{"g", "o", "i", "n", "g"},
			{"'", "", "", "", ""},
			{"o", "n", "", "", ""},
			{"h", "e", "r", "e", ""},
			{"w", "t", "f", "", ""},
		}, tokens)
	})
}
