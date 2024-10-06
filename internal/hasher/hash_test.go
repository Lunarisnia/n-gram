package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hash(t *testing.T) {
	t.Run("Expect properly hashed", func(t *testing.T) {
		hashMap := NewHashMap()
		tokens := [][]string{
			{"c", "a", "r"},
			{"c", "a", "t"},
			{"c", "a", "r"},
		}
		hashMap.InsertNode(tokens[0])
		hashMap.InsertNode(tokens[1])
		hashMap.InsertNode(tokens[2])
		assert.Equal(t, 2, hashMap.Data[310].Counter)
		assert.Equal(t, 1, hashMap.Data[312].Counter)
	})
}
