package tree

import "errors"

// BinaryTree 二叉树
type BinaryTree struct {
	*Node
}

// NewBinaryTree 新建二叉数
func NewBinaryTree(val interface{}) *BinaryTree {
	return &BinaryTree{
		Node: NewNode(val),
	}
}

// AddChild 重载函数
func (bt *BinaryTree) AddChild(child *BinaryTree) error {
	if len(bt.children) == 2 {
		return errors.New("already has 2 children")
	}
	return bt.Node.AddChild(child.Node)
}
