package tree

import "errors"

// BinaryNode 二叉树
type BinaryNode struct {
	*Node
}

// NewBinaryNode 新建二叉树
func NewBinaryNode(val interface{}) *BinaryNode {
	return &BinaryNode{
		Node: NewNode(val),
	}
}

// AddChild 重载函数
func (bt *BinaryNode) AddChild(child TNode) error {
	c := child.(*BinaryNode)
	if bt.Degree() == 2 {
		return errors.New("already has 2 children")
	}
	return bt.AddChild(c)
}

// BinaryThrNode 线索二叉树
type BinaryThrNode struct {
	*Node
	lchild TNode
	rchild TNode
}

// AddChild 重载
func (bt *BinaryThrNode) AddChild(child TNode) error {
	c := child.(*BinaryThrNode)
	dg := bt.Degree()
	if dg == 0 {
		bt.lchild = c
	} else if dg == 1 {
		bt.rchild = c
	} else {
		return errors.New("already has 2 children")
	}

	return bt.AddChild(c)
}
