package list

import "testing"

func TestReverse(t *testing.T) {
	ll := &LinkedList{Val: 1, Next: &LinkedList{Val: 2, Next: &LinkedList{Val: 3, Next: &LinkedList{Val: 4, Next: &LinkedList{Val: 5, Next: &LinkedList{Val: 6, Next: nil}}}}}}
	ll2 := ll.Reverse()
	t.Log(ll2)
}
