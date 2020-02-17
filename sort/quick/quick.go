package main

import (
	"algorithmGo/sort"
	"log"
	"math/rand"
)

func main() {
	for i := 0; i < sort.TestTime; i++ {
		arr := sort.GenerateRandomArray(sort.Size, sort.Value)
		brr := sort.CopyArr(arr)
		quickSort(arr)
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

func quickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	quickSortImpl(arr, 0, len(arr)-1)
}

func quickSortImpl(arr []int, left, right int) {
	if left < right {
		tmpIndex := left + rand.Intn(right-left+1)
		arr[right], arr[tmpIndex] = arr[tmpIndex], arr[right]
		midL, midR := partition(arr, left, right)
		quickSortImpl(arr, left, midL-1)
		quickSortImpl(arr, midR+1, right)
	}
}

func partition(arr []int, left, right int) (int, int) {
	less := left - 1
	more := right

	for left < more {
		if arr[left] < arr[right] {
			less++
			arr[less], arr[left] = arr[left], arr[less]
			left++
		} else if arr[left] > arr[right] {
			more--
			arr[left], arr[more] = arr[more], arr[left]
		} else {
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	return less + 1, more
}
