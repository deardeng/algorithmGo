package main

import (
	"algorithmGo/list/203"
	"log"
)

func removeElements(head *_03.ListNode, val int) *_03.ListNode {
	if head == nil {
		return nil
	}

	res := removeElements(head.Next, val)
	if head.Val == val {
		return res
	} else {
		head.Next = res
		return head
	}
}

func main() {
	a := []int{1, 2, 6, 3, 4, 5, 6}
	list := &_03.ListNode{}
	list.NewListNode(a)
	log.Print("before: ", list)
	list = removeElements(list, 6)
	log.Print("after: ", list)
}
