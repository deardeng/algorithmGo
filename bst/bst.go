package bst

import "log"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
	size int
}

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

//// 向二分搜索树中添加新的元素e
//func (bst *BST) Add(val int) {
//	if bst.root == nil {
//		bst.root = &Node{}
//		bst.size++
//	} else {
//		bst.add(bst.root, val)
//	}
//}
//
//// 向以node为根的二分搜索树中插入元素E，递归算法
//func (bst *BST) add(node *Node, e int) {
//	if e == node.value {
//		return
//	} else if e < node.value && node.left == nil {
//		node.left = &Node{value: e}
//		bst.size++
//		return
//	} else if e > node.value && node.right == nil {
//		node.right = &Node{value: e}
//		bst.size++
//		return
//	}
//
//	if e < node.value {
//		bst.add(node.left, e)
//	} else { // e > node.value
//		bst.add(node.right, e)
//	}
//}

// 向二分搜索树中添加新的元素e
func (bst *BST) Add(val int) {
	bst.root = bst.add(bst.root, val)
}

// 向以node为根的二分搜索树中插入元素E，递归算法
// 返回插入新节点后二分搜索树的根
func (bst *BST) add(node *Node, e int) *Node {
	if node == nil {
		bst.size++
		return &Node{value: e}
	}

	if e < node.value {
		node.left = bst.add(node.left, e)
	} else if e > node.value {
		node.right = bst.add(node.right, e)
	}

	return node
}

// 看二分搜索树中是否包含元素e
func (bst *BST) Contains(e int) bool {
	return bst.contains(bst.root, e)
}

// 看以node为根的二分搜索树中是否包含元素e，递归算法
func (bst *BST) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}

	if e == node.value {
		return true
	} else if e < node.value {
		return bst.contains(node.left, e)
	} else {
		return bst.contains(node.right, e)
	}
}

// 二分搜索树的前序遍历
func (bst *BST) PreOrder() {
	bst.preOrder(bst.root)
}

// 前序遍历以node为根的二分搜索树，递归算法
func (bst *BST) preOrder(node *Node) {
	if node == nil {
		return
	}

	log.Print(node.value)
	bst.preOrder(node.left)
	bst.preOrder(node.right)
}
