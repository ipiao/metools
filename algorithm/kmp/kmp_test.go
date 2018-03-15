package kmp

import "testing"

func TestNext(t *testing.T) {
	next := getNext([]byte("abaabc"))
	t.Log(next)
	ind := StringIndex("adaskbk", "bk")
	t.Log(ind)
}
