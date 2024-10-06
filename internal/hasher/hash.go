package hasher

import "slices"

type TokenNode struct {
	Token   []string
	Counter int
}

func NewTokenNode(token []string) *TokenNode {
	return &TokenNode{
		Token:   token,
		Counter: 1,
	}
}

type HashMap struct {
	Data map[int]*TokenNode
}

func convertToKey(token []string) int {
	key := 0
	for _, ch := range token {
		r := rune(ch[0])
		key += int(r)
	}
	return key
}

func NewHashMap() *HashMap {
	return &HashMap{
		Data: make(map[int]*TokenNode),
	}
}

func (h *HashMap) InsertNode(token []string) {
	key := convertToKey(token)
	newNode := NewTokenNode(token)
	if node, exist := h.Data[key]; exist {
		if slices.Equal(node.Token, token) {
			node.Counter++
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
