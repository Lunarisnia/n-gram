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
			{"h", "e", "y", "_", "_"},
			{"t", "H", "e", "r", "e"},
			{"w", "h", "a", "t", "_"},
			{"i", "s", "_", "_", "_"},
			{"g", "o", "i", "n", "g"},
			{"'", "_", "_", "_", "_"},
			{"o", "n", "_", "_", "_"},
			{"h", "e", "r", "e", "_"},
			{"w", "t", "f", "_", "_"},
		}, tokens)
	})
}
