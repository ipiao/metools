package tree

import (
	"log"
	"sort"
)

// WNode 带权值的节点
type WNode struct {
	*Node
	weight int
}

// Huffman 由所有的权节点生成huffman树
// 1.获取权值最小的两棵树作为左右子树构建一棵新的二叉树，且置根节点的权值为2棵子树的和
// 2.删除两棵子树，
func Huffman(nodes []WNode) *WNode {
	length := len(nodes)
	if length == 0 {
		return nil
	} else if length == 1 {
		return &nodes[0]
	} else if length == 2 {
		nNode := &WNode{Node: NewNode(nil), weight: nodes[0].weight + nodes[1].weight}
		nNode.AddChildren(&nodes[0], &nodes[1])
		return nNode
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].weight < nodes[j].weight
	})
Out:
	for length > 2 {
		length--
		log.Println("len nodes is ", length)
		nNode := Huffman(nodes[:2])
		nnodes := make([]WNode, length)
		for i := 2; i < len(nodes); i++ {
			if nodes[i].weight >= nNode.weight {
				copy(nnodes[:i-2], nodes[2:i])
				nnodes[i-2] = *nNode
				copy(nnodes[i-1:], nodes[i:])
				nodes = nnodes
				goto Out
			}
		}
		nodes = append(nodes[2:], *nNode)
	}
	return Huffman(nodes)
}
