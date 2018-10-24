package list

import "encoding/json"

// LinkedList 链表,单链表
type LinkedList struct {
	Val  interface{}
	Next *LinkedList
}

// Reverse 链表反转
func (l *LinkedList) Reverse() *LinkedList {
	return reverse(l)
}

func reverse(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}
	var pHead, next *LinkedList
	pHead = head
	head = head.Next
	pHead.Next = nil
	next = head.Next
	for head != nil {
		head.Next = pHead
		pHead = head
		head = next
		if next != nil {
			next = next.Next
		}
	}
	return pHead
}
func (l *LinkedList) String() string {
	bs, _ := json.Marshal(l)
	return string(bs)
}

func reverse2(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverse2(head.Next)
	if head.Next != nil {
		head.Next.Next = head
		head.Next = nil
	}
	return node
}
