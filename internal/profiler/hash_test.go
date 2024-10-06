package profiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hash(t *testing.T) {
	t.Run("Expect properly hashed", func(t *testing.T) {
		hashMap := NewHashMap()
		hashMap.InsertNode("car")
		hashMap.InsertNode("cat")
		hashMap.InsertNode("car")
		assert.Equal(t, 2, hashMap.Data[310].Counter)
		assert.Equal(t, 1, hashMap.Data[312].Counter)
	})

	t.Run("Expect rank properly calculated", func(t *testing.T) {
		hashMap := NewHashMap()
		hashMap.InsertNode("car")
		hashMap.InsertNode("cat")
		hashMap.InsertNode("car")

		hashMap.CalculateAllRanks()
		assert.Equal(t, 0, hashMap.Data[310].Rank)
		assert.Equal(t, 1, hashMap.Data[312].Rank)
	})
}
