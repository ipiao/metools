package tree

import "errors"

// TNode 树根，树根是没有向上的
type TNode interface {
	SetVal(val interface{})
	AddChild(tree Tree) error
	WhichChild(Tree) int
	DeleteChild(i int)
	Child(i int) (Tree, bool)
	Children() []Tree
	Degree() int
	Tree
}

// Tree 子树
type Tree interface {
	Parent() TNode //  父树根
	SetParent(p TNode)
	Root() TNode
	WhichChildOfParent() int
	LevelChildren(lev int) []Tree
	// Depth() int
}

var tnode TNode = &node{}
var ttree Tree = &node{}

//*******************

// TNode 定义树结构
// type TNode interface {
//
// 	Depth() int
// 	Level() int

// 	GenerationsNum() int

// 	LeftSibling() (Tree, bool)
// 	RightSibling() (Tree, bool)

// 	LeafGenerations() []Tree
//

// 	Map() map[string]interface{}
// }

type node struct {
	val      interface{}
	children []Tree
	parent   TNode
}

func (n *node) SetVal(val interface{}) {
	n.val = val
}

func (n *node) AddChild(child Tree) error {
	if child != nil && child.Parent() != nil {
		return errors.New("child has parent")
	}
	child.SetParent(n)
	n.children = append(n.children, child)
	return nil
}

func (n *node) DeleteChild(i int) {
	if i < len(n.children) && i >= 0 {
		n.children[i].SetParent(nil)
		n.children = append(n.children[:i], n.children[i+1:]...)
	}
}

func (n *node) WhichChild(child Tree) int {
	for i, c := range n.children {
		if c == child {
			return i
		}
	}
	return -1
}

func (n *node) Child(i int) (Tree, bool) {
	if i >= len(n.children) || 0 < i {
		return nil, false
	}
	return n.children[i], true
}

func (n *node) Children() []Tree {
	return n.children
}

func (n *node) Degree() int {
	return len(n.children)
}

func (n *node) Parent() TNode {
	return n.parent
}

func (n *node) SetParent(p TNode) {
	n.parent = p
}

func (n *node) Root() TNode {
	if n.parent == nil {
		return n
	}
	root := n.Parent()
	for root.Parent() != nil {
		root = root.Parent()
	}
	return root
}

func (n *node) WhichChildOfParent() int {
	if n.parent != nil {
		return n.parent.WhichChild(n)
	}
	return -1
}

func (n *node) LevelChildren(lev int) []Tree {
	ret := make([]Tree, 0)
	if lev == 1 {
		ret = append(ret, n.children...)
	} else if lev > 1 {
		for _, c := range n.children {
			ret = append(ret, c.LevelChildren(lev-1)...)
		}
	}
	return ret
}

// func (n *node) Depth() int {
// 	lev := 1
// 	for len(n.LevelChildren(lev)) != 0 {
// 		lev++
// 	}
// 	return lev
// }
