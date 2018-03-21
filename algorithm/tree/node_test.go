package tree

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	n0 := NewNode("-")
	n10 := NewNode("+")
	n11 := NewNode("/")
	n0.AddChildren(n10, n11)
	n20 := NewNode("a")
	n21 := NewNode("*")
	n22 := NewNode("e")
	n23 := NewNode("f")
	n10.AddChildren(n20, n21)
	n11.AddChildren(n22, n23)
	n30 := NewNode("b")
	n31 := NewNode("-")
	n21.AddChildren(n30, n31)
	n40 := NewNode("c")
	n41 := NewNode("d")
	n31.AddChildren(n40, n41)
	t.Log(n0.Map())
	n0.Root().DLRVisit(func(n TNode) error {
		fmt.Print(n.GetVal())
		return nil
	})
	fmt.Println()
	n0.Root().LDRVisit(func(n TNode) error {
		fmt.Print(n.GetVal())
		return nil
	})
	fmt.Println()
	n0.Root().LRDVisit(func(n TNode) error {
		fmt.Print(n.GetVal())
		return nil
	})
}

func TestBT(t *testing.T) {
	n0 := NewNode("1")
	n10 := NewNode("2")
	n101 := NewNode("21")
	n10.AddChild(n101)
	n11 := NewNode("3")
	n111 := NewNode("31")
	n11.AddChild(n111)
	n12 := NewNode("4")
	n121 := NewNode("41")
	n12.AddChild(n121)
	// tns := []TNode{n10, n11, n12}
	// bt := bforest2BTree(tns)
	n0.AddChildren(n10, n11, n12)
	bt := n0.BinaryTree()
	bs, _ := json.Marshal(bt.Map())
	t.Log(string(bs))
}

func TestBTT(t *testing.T) {
	n10 := NewNode("2")
	n101 := NewNode("21")
	n10.AddChild(n101)
	n11 := NewNode("3")
	n111 := NewNode("31")
	n11.AddChild(n111)
	n12 := NewNode("4")
	n121 := NewNode("41")
	n12.AddChild(n121)
	tns := []TNode{n10, n11, n12}
	bt := bforest2BTree(tns)

	bs, _ := json.Marshal(bt.Map())
	t.Log(string(bs))
}
