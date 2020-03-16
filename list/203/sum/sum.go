package main

import "log"

func sum(arr []int) int {
	return sumImpl(arr, 0)
}

// 计算arr[l ... n)这个区间内所有数字的和
func sumImpl(arr []int, l int) int {
	if l == len(arr) {
		return 0
	}
	return arr[l] + sumImpl(arr, l+1)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	log.Print(sum(nums))
}
