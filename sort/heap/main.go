package main

import (
	"algorithmGo/sort"
	"log"
)

func main() {
	for i := 0; i < sort.TestTime; i++ {
		arr := sort.GenerateRandomArray(sort.Size, sort.Value)
		brr := sort.CopyArr(arr)
		heapSort(arr)
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

func heapInsert(array []int, index int) {
	for array[index] > array[(index-1)/2] {
		array[index], array[(index-1)/2] = array[(index-1)/2], array[index]
		index = (index - 1) / 2
	}
}

func heapify(array []int, index int, size int) {
	left := index*2 + 1
	for left < size {
		var largest int
		if left+1 < size && array[left+1] > array[left] {
			largest = left + 1
		} else {
			largest = left
		}

		if array[largest] < array[index] {
			largest = index
		}

		if largest == index {
			break
		}
		array[largest], array[index] = array[index], array[largest]
		index = largest
		left = index*2 + 1
	}
}

func heapSort(array []int) {
	if array == nil || len(array) < 2 {
		return
	}

	for i := 0; i < len(array); i++ {
		heapInsert(array, i)
	}

	size := len(array)
	size--
	array[0], array[size] = array[size], array[0]
	for size > 0 {
		heapify(array, 0, size)
		size--
		array[0], array[size] = array[size], array[0]
	}
}
