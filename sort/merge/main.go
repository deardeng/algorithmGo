package main

import (
	"algorithmGo/sort"
	"log"
)

func mergeSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	sortProcess(arr, 0, len(arr)-1)
}

func sortProcess(arr []int, left, right int) {
	if left == right {
		return
	}

	mid := left + (right-left)>>1
	sortProcess(arr, left, mid)
	sortProcess(arr, mid+1, right)

	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	helper := make([]int, right-left+1)
	i := 0
	p1, p2 := left, mid+1
	for p1 <= mid && p2 <= right {
		if arr[p1] < arr[p2] {
			helper[i] = arr[p1]
			p1++
		} else {
			helper[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= mid {
		helper[i] = arr[p1]
		p1++
		i++
	}

	for p2 <= right {
		helper[i] = arr[p2]
		p2++
		i++
	}

	for i := 0; i < len(helper); i++ {
		arr[left+i] = helper[i]
	}
}

func main() {
	for i := 0; i < sort.TestTime; i++ {
		arr := sort.GenerateRandomArray(sort.Size, sort.Value)
		brr := sort.CopyArr(arr)
		mergeSort(arr)
		sort.RightMethod(brr)
		if sort.IsEqual(arr, brr) == false {
			sort.Succeed = false
			break
		}
	}
	if sort.Succeed != true {
		log.Printf("bad")
		return
	}
	log.Printf("good")
}
