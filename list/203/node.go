package _03

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表节点的构造函数
// 使用arr为参数，创建一个链表，当前的ListNode为链表头节点
func (ln *ListNode) NewListNode(array []int) {
	if array == nil || len(array) == 0 {
		panic("arr can not be empty")
	}

	ln.Val = array[0]

	cur := ln
	for i := 1; i < len(array); i++ {
		cur.Next = &ListNode{
			Val: array[i],
		}
		cur = cur.Next
	}
}

// 以当前节点为头节点的链表信息字符串
func (ln *ListNode) String() string {
	result := ""
	cur := ln
	for cur != nil {
		result += strconv.Itoa(cur.Val) + "->"
		cur = cur.Next
	}
	result += "nil"
	return result
}
