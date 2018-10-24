package tree

import (
	"errors"
	"log"
)

// k 树的性质
// 节点数 n = n0 + n1 + n2 + ... + nk(nk 是具有k个子节点的节点数)
// n = n1 + 2n2 +3n3 + ... + knk
// 所以 n0 = n2 + 2n3 + ... + (k-1)nk

// TNode 树根，树根是没有向上的
type TNode interface {
	SetVal(val interface{})
	GetVal() interface{}
	Tree
}

// Tree 子树
type Tree interface {
	AddChild(TNode) error
	WhichChild(TNode) int
	DeleteChild(i int)
	Child(i int) (TNode, bool)
	Children() []TNode
	LevelChildren(lev int) []TNode
	LeafGenerations() []TNode

	Degree() int
	Parent() TNode //  父树根
	SetParent(p TNode)
	Root() TNode
	WhichChildOfParent() int

	Depth() int
	GenerationsNum() int

	LeftSibling() (TNode, bool)
	RightSibling() (TNode, bool)
	LeftChild() (TNode, bool)
	RightChild() (TNode, bool)

	DLRVisit(fn func(TNode) error) error
	LDRVisit(fn func(n TNode) error) error
	LRDVisit(fn func(n TNode) error) error

	Map() map[string]interface{}

	BinaryTree() TNode
	Copy() TNode
}

// Node 节点
type Node struct {
	val      interface{}
	children []TNode
	parent   TNode
}

// NewNode 新建节点
func NewNode(val interface{}) *Node {
	return &Node{
		val: val,
	}
}

// GetVal 获取节点值
func (n *Node) GetVal() interface{} {
	return n.val
}

// SetVal 设置节点值
func (n *Node) SetVal(val interface{}) {
	n.val = val
}

// AddChild 添加子节点
func (n *Node) AddChild(child TNode) error {
	if child != nil && child.Parent() != nil {
		return errors.New("child has parent")
	}
	child.SetParent(n)
	n.children = append(n.children, child)
	return nil
}

// AddChildren 添加子节点
func (n *Node) AddChildren(children ...TNode) error {
	for _, child := range children {
		if err := n.AddChild(child); err != nil {
			return err
		}
	}
	return nil
}

// DeleteChild 删除子节点
func (n *Node) DeleteChild(i int) {
	if i < len(n.children) && i >= 0 {
		n.children[i].SetParent(nil)
		n.children = append(n.children[:i], n.children[i+1:]...)
	}
}

// WhichChild 第几个子节点
func (n *Node) WhichChild(child TNode) int {
	for i, c := range n.children {
		if c == child {
			return i
		}
	}
	return -1
}

// Child 获取子节点
func (n *Node) Child(i int) (TNode, bool) {
	if i >= len(n.children) || 0 < i {
		return nil, false
	}
	return n.children[i], true
}

// Children 获取所有子节点
func (n *Node) Children() []TNode {
	return n.children
}

// Degree 度
func (n *Node) Degree() int {
	return len(n.children)
}

// Parent 父节点
func (n *Node) Parent() TNode {
	return n.parent
}

// SetParent 设置父节点
func (n *Node) SetParent(p TNode) {
	n.parent = p
}

// Root 所在树的根节点
func (n *Node) Root() TNode {
	if n.parent == nil {
		return n
	}
	root := n.Parent()
	for root.Parent() != nil {
		root = root.Parent()
	}
	return root
}

// WhichChildOfParent 节点属于父节点的哪个子节点
func (n *Node) WhichChildOfParent() int {
	if n.parent != nil {
		return n.parent.WhichChild(n)
	}
	return -1
}

// LevelChildren 当前节点地往下第i层的所有子节点
func (n *Node) LevelChildren(lev int) []TNode {
	ret := make([]TNode, 0)
	if lev == 1 {
		ret = append(ret, n.children...)
	} else if lev > 1 {
		for _, c := range n.children {
			ret = append(ret, c.LevelChildren(lev-1)...)
		}
	}
	return ret
}

// Depth 当前节点为根节点的树的深度
func (n *Node) Depth() int {
	lev := 1
	for len(n.LevelChildren(lev)) != 0 {
		lev++
	}
	return lev
}

// Level 当前节点在所在树的第几层
func (n *Node) Level() int {
	lev := 0
	var p = n.Parent()
	for p != nil {
		lev++
		p = p.Parent()
	}
	return lev
}

// GenerationsNum 当前节点所有子孙后代的个数
func (n *Node) GenerationsNum() int {
	num := len(n.children)
	for _, c := range n.children {
		num += c.GenerationsNum()
	}
	return num
}

// LeftSibling 当前节点的左兄弟节点
func (n *Node) LeftSibling() (TNode, bool) {
	index := n.WhichChildOfParent()
	if index <= 0 {
		return nil, false
	}
	return n.parent.Child(index - 1)
}

// RightSibling 当前节点的右兄弟节点
func (n *Node) RightSibling() (TNode, bool) {
	index := n.WhichChildOfParent()
	if index == -1 {
		return nil, false
	}
	if n.parent.Degree() < index+2 {
		return nil, false
	}
	return n.parent.Child(index + 1)
}

// LeftChild 左孩子
func (n *Node) LeftChild() (TNode, bool) {
	if len(n.children) > 0 {
		return n.children[0], true
	}
	return nil, false
}

// RightChild 右孩子
func (n *Node) RightChild() (TNode, bool) {
	l := len(n.children)
	if l > 0 {
		return n.children[l-1], true
	}
	return nil, false
}

// LeafGenerations 当前节点所有的叶子节点后代
func (n *Node) LeafGenerations() []TNode {
	leafs := make([]TNode, 0)
	for _, c := range n.children {
		if c.Degree() == 0 {
			leafs = append(leafs, c)
		} else {
			leafs = append(leafs, c.LeafGenerations()...)
		}
	}
	return leafs
}

// Map 转换成map形式
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

// DLRVisit 先(根)序遍历
func (n *Node) DLRVisit(fn func(TNode) error) error {
	err := fn(n)
	if err != nil {
		return err
	}
	lnode, has := n.LeftChild()
	if has {
		err = lnode.DLRVisit(fn)
		if err != nil {
			return err
		}
	}
	rnode, has := n.RightChild()
	if has {
		err = rnode.DLRVisit(fn)
		if err != nil {
			return err
		}
	}
	return nil
}

// LDRVisit 中(根)序遍历
func (n *Node) LDRVisit(fn func(TNode) error) error {
	var err error
	lnode, has := n.LeftChild()
	if has {
		err = lnode.LDRVisit(fn)
		if err != nil {
			return err
		}
	}
	err = fn(n)
	if err != nil {
		return err
	}
	rnode, has := n.RightChild()
	if has {
		err = rnode.LDRVisit(fn)
		if err != nil {
			return err
		}
	}
	return err
}

// LRDVisit 后(根)序遍历
func (n *Node) LRDVisit(fn func(TNode) error) error {
	var err error
	lnode, has := n.LeftChild()
	if has {
		err = lnode.LRDVisit(fn)
		if err != nil {
			return err
		}
	}
	rnode, has := n.RightChild()
	if has {
		err = rnode.LRDVisit(fn)
		if err != nil {
			return err
		}
	}
	err = fn(n)
	if err != nil {
		return err
	}
	return err
}

// BinaryTree 树生成二叉树
func (n *Node) BinaryTree() TNode {
	nn := NewNode(n.val)
	dg := n.Degree()
	if dg > 0 {
		cls := make([]TNode, dg)
		for i := 0; i < dg; i++ {
			cp := n.children[i]
			log.Println(cp.Map())
			cls[i] = cp.BinaryTree()
		}
		nn.AddChild(bforest2BTree(cls))
	}
	return nn
}

// Copy 复制
func (n *Node) Copy() TNode {
	var nn = NewNode(n.val)
	for i := range n.children {
		nn.AddChild(n.children[i].Copy())
	}
	return nn
}

func bforest2BTree(nodes []TNode) TNode {
	if len(nodes) == 1 {
		return nodes[0]
	} else if len(nodes) == 2 {
		fnode := nodes[0]
		enode := nodes[1]
		fnode.AddChild(enode)
		return fnode
	}
	fnod := bforest2BTree(nodes[1:])
	return bforest2BTree([]TNode{nodes[0], fnod})
}

// LevelVist 层次遍历
func (n *Node) LevelVist(fn func(TNode) error) error {
	// 1.首先将根节点放入队列中。
	// 2.当队列为非空时，循环执行步骤3到步骤5，否则执行6；
	// 3.出队列取得一个结点，访问该结点；
	// 4.若该结点的左子树为非空，则将该结点的左子树入队列；
	// 5.若该结点的右子树为非空，则将该结点的右子树入队列；
	// 6.结束。
	return nil
}
