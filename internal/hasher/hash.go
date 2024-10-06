package hasher

import "slices"

type TokenNode struct {
	Token   []string
	Counter int
	Next    *TokenNode
}

func NewTokenNode(token []string) *TokenNode {
	return &TokenNode{
		Token:   token,
		Counter: 1,
		Next:    nil,
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

func (h *HashMap) appendNode(parentNode *TokenNode, newNode *TokenNode) {
	if slices.Equal(parentNode.Token, newNode.Token) {
		parentNode.Counter++
		return
	}
	if parentNode.Next == nil {
		parentNode.Next = newNode
		return
	}

	h.appendNode(parentNode.Next, newNode)
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
		h.appendNode(node, newNode)
	} else {
		h.Data[key] = newNode
	}
}
