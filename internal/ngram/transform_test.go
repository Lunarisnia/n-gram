package ngram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeNGram(t *testing.T) {
	t.Run("Test encode a bigram", func(t *testing.T) {
		result, err := Encode("TEXT", 2)
		assert.Nil(t, err)
		fmt.Println("Result: ", result)
	})
	t.Run("Test encode a trigram", func(t *testing.T) {
		result, err := Encode("TEXT", 3)
		assert.Nil(t, err)
		fmt.Println("Result: ", result)
	})
	t.Run("Test encode a quadgram", func(t *testing.T) {
		result, err := Encode("TEXT", 4)
		assert.Nil(t, err)
		fmt.Println("Result: ", result)
	})
}
