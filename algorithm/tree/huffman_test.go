package tree

import (
	"encoding/json"
	"testing"
)

func TestHuffman(t *testing.T) {
	wnodes := []WNode{
		{Node: NewNode(1), weight: 1},
		{Node: NewNode(2), weight: 2},
		{Node: NewNode(3), weight: 3},
		{Node: NewNode(4), weight: 4},
		{Node: NewNode(5), weight: 5},
		{Node: NewNode(6), weight: 6},
		{Node: NewNode(7), weight: 7},
		{Node: NewNode(8), weight: 8},
	}
	nnode := Huffman(wnodes)
	bs, _ := json.Marshal(nnode.Map())
	t.Log(string(bs))
}
