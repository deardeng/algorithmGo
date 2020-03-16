package main

import "algorithmGo/bst"

func main() {
	bst := bst.BST{}
	nums := []int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}
	bst.PreOrder()
}
