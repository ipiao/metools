package tree

import (
	"errors"
)

// Node 节点
type Node struct {
	val      interface{}
	children []*Node
	parent   *Node
}

// Map 构造马屁格式
func (n *Node) Map() map[string]interface{} {
	m := make(map[string]interface{}, 2)
	m["val"] = n.val
	chis := make([]map[string]interface{}, len(n.children))
	for i, c := range n.children {
		chis[i] = c.Map()
	}
	m["children"] = chis
	return m
}

// NewNode 新建一个节点
func NewNode(val interface{}) *Node {
	return &Node{
		val:      val,
		children: make([]*Node, 0),
		parent:   nil,
	}
}

// SetVal 设置节点值
func (n *Node) SetVal(val interface{}) {
	n.val = val
}

// AddChild 添加子节点
func (n *Node) AddChild(child *Node) error {
	if child != nil && child.parent != nil {
		return errors.New("child has parent")
	}
	child.parent = n
	n.children = append(n.children, child)
	return nil
}

// Parent 父节点
func (n *Node) Parent() *Node {
	return n.parent
}

// Degree 度，拥有字节点的个数
func (n *Node) Degree() int {
	return len(n.children)
}

// Child 获取第(i+1)个子节点
func (n *Node) Child(i int) (*Node, bool) {
	if i >= len(n.children) || 0 < i {
		return nil, false
	}
	return n.children[i], true
}

// DeleteChild 删除第(i+1)个子节点
func (n *Node) DeleteChild(i int) {
	if i < len(n.children) && i >= 0 {
		n.children[i].parent = nil
		n.children = append(n.children[:i], n.children[i+1:]...)
	}
}

// HasChild 验证是否有字节点，防伪造
func (n *Node) HasChild(child *Node) bool {
	for _, c := range n.children {
		if c == child {
			return true
		}
	}
	return false
}

// WhichChild child 属于第几个孩子
func (n *Node) WhichChild(child *Node) int {
	for i, c := range n.children {
		if c == child {
			return i
		}
	}
	return -1
}

// LeftChild 返回左子节点
func (n *Node) LeftChild() (*Node, bool) {
	if len(n.children) > 0 {
		return n.children[0], true
	}
	return nil, false
}

// RightChild 返回右子节点
func (n *Node) RightChild() (*Node, bool) {
	l := len(n.children)
	if l > 0 {
		return n.children[l-1], true
	}
	return nil, false
}

// IsRoot 是否是根节点
func (n *Node) IsRoot() bool {
	return n.parent == nil
}

// WhichChildOfParent 属于父节点的第(i-1)个字节点
func (n *Node) WhichChildOfParent() int {
	if n.parent != nil {
		return n.parent.WhichChild(n)
	}
	return -1
}

// LeftSibling 左兄弟节点
func (n *Node) LeftSibling() (*Node, bool) {
	index := n.WhichChildOfParent()
	if index <= 0 {
		return nil, false
	}
	return n.parent.children[index-1], true
}

// RightSibling 右兄弟节点
func (n *Node) RightSibling() (*Node, bool) {
	index := n.WhichChildOfParent()
	if index == -1 {
		return nil, false
	}
	if n.parent.Degree() < index+2 {
		return nil, false
	}
	return n.parent.children[index+1], true
}

// Level 节点所在树的层次，根节点层次为0
func (n *Node) Level() int {
	lev := 0
	var p = n
	for p.parent != nil {
		lev++
		p = p.parent
	}
	return lev
}

// IsLeaf 是否是叶节点
func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

// Root 返回节点所在树的根节点
func (n *Node) Root() *Node {
	var root = n
	for root.parent != nil {
		root = root.parent
	}
	return root
}

// GenerationsNum 节点所有子孙数
func (n *Node) GenerationsNum() int {
	num := len(n.children)
	for _, c := range n.children {
		num += c.GenerationsNum()
	}
	return num
}

// TreeRootGenerationsNum 节点所在树的根节点的子孙数
func (n *Node) TreeRootGenerationsNum() int {
	root := n.Root()
	return root.GenerationsNum()
}

// LevelChildren 节点层次 lev 的字节点,lev>1
func (n *Node) LevelChildren(lev int) []*Node {
	ret := make([]*Node, 0)
	if lev == 1 {
		ret = append(ret, n.children...)
	} else if lev > 1 {
		for _, c := range n.LevelChildren(lev - 1) {
			ret = append(ret, c.children...)
		}
	}
	return ret
}

// TreeRootLevelChildren 节点所在树的根节点层次 lev 的字节点,lev>1
func (n *Node) TreeRootLevelChildren(lev int) []*Node {
	root := n.Root()
	return root.LevelChildren(lev)
}

// Depth 深度
func (n *Node) Depth() int {
	lev := 1
	for len(n.LevelChildren(lev)) != 0 {
		lev++
	}
	return lev
}

// TreeDepth 节点所在树的深度
func (n *Node) TreeDepth() int {
	root := n.Root()
	return root.Depth()
}

// LeafGenerations 节点所叶节点子孙
func (n *Node) LeafGenerations() []*Node {
	leafs := make([]*Node, 0)
	for _, c := range n.children {
		if len(c.children) == 0 {
			leafs = append(leafs, c)
		} else {
			leafs = append(leafs, c.LeafGenerations()...)
		}
	}
	return leafs
}

// TreeLeafGenerations 所在树的所有叶节点
func (n *Node) TreeLeafGenerations() []*Node {
	root := n.Root()
	return root.LeafGenerations()
}
