package main

import (
	"algorithmGo/list/203"
	"log"
)

func removeElements(head *_03.ListNode, val int) *_03.ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return head
}

func main() {
	a := []int{1, 2, 6, 3, 4, 5, 6}
	list := &_03.ListNode{}
	list.NewListNode(a)
	log.Print("before: ", list)
	list = removeElements(list, 6)
	log.Print("after: ", list)
}
