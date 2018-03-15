package tree

import (
	"testing"
)

func TestNode(t *testing.T) {
	n0 := NewNode("hello")
	n01 := NewNode("hello1")
	n0.AddChild(n01)
	t.Log(n0.Map())
	t.Log(n0, n01)
	t.Log(n01.Level())
	t.Log(n0.Degree(), n0.Depth())
	n0.DeleteChild(0)
	t.Log(n0, n01)
}
