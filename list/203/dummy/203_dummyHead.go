package main

import (
	"algorithmGo/list/203"
	"log"
)

func removeElementsDummyHead(head *_03.ListNode, val int) *_03.ListNode {
	dummyHead := &_03.ListNode{
		-1,
		head,
	}

	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return dummyHead.Next
}

func main() {
	a := []int{1, 2, 6, 3, 4, 5, 6}
	list := &_03.ListNode{}
	list.NewListNode(a)
	log.Print("before: ", list)
	list = removeElementsDummyHead(list, 6)
	log.Print("after: ", list)
}
