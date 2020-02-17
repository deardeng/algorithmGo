package main

import (
	"algorithmGo/sort"
	"log"
)

func main() {
	for i := 0; i < sort.TestTime; i++ {
		arr := sort.GenerateRandomArray(sort.Size, sort.Value)
		brr := sort.CopyArr(arr)
		insertSort(arr)
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

func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}
