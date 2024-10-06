package profiler

import (
	"sort"
)

type TokenNode struct {
	Token   string
	Counter int
	Rank    int
}

func NewTokenNode(token string) *TokenNode {
	return &TokenNode{
		Token:   token,
		Counter: 1,
		Rank:    0,
	}
}

type HashMap struct {
	Data map[int]*TokenNode
}

func convertToKey(token string) int {
	key := 0
	for _, r := range token {
		key += int(r)
	}
	return key
}

func NewHashMap() *HashMap {
	return &HashMap{
		Data: make(map[int]*TokenNode),
	}
}

func (h *HashMap) InsertNode(token string) {
	key := convertToKey(token)
	newNode := NewTokenNode(token)
	if node, exist := h.Data[key]; exist {
		if node.Token == token {
			node.Counter++
			return
		}
		for {
			key++
			if _, exist := h.Data[key]; !exist {
				h.Data[key] = newNode
				return
			}
		}
	} else {
		h.Data[key] = newNode
	}
}

func (h *HashMap) CalculateAllRanks() {
	tokenNodes := make([]*TokenNode, 0)
	for _, node := range h.Data {
		tokenNodes = append(tokenNodes, node)
	}
	sort.Slice(tokenNodes, func(i, j int) bool {
		return tokenNodes[i].Counter > tokenNodes[j].Counter
	})
	for i, node := range tokenNodes {
		node.Rank = i
	}
}
