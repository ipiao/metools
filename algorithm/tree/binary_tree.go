package tree

// // 二叉树的一个性质

// // BinaryTree 二叉树
// type BinaryTree Node

// // NewBinaryTree 新建二叉数
// func NewBinaryTree(val interface{}) *BinaryTree {
// 	return &BinaryTree{
// 		val:      val,
// 		children: make([]*Node, 0),
// 		parent:   nil,
// 	}
// }

// // AddChild 重载函数
// func (bt *BinaryTree) AddChild(child *BinaryTree) error {
// 	if len(bt.children) == 2 {
// 		return errors.New("already has 2 children")
// 	}
// 	if child != nil && child.parent != nil {
// 		return errors.New("child has parent")
// 	}
// 	child.parent = bt
// 	n.children = append(n.children, child)
// 	return nil
// }
