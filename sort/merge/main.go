package main

import (
	"algorithmGo/sort"
	"algorithmGo/sort/merge/smallSum"
	"log"
)

func main() {
	// 归并排序
	for i := 0; i < sort.TestTime; i++ {
		arr := sort.GenerateRandomArray(sort.Size, sort.Value)
		brr := sort.CopyArr(arr)
		smallSum.MergeSort(arr)
		sort.RightMethod(brr)
		if sort.IsEqual(arr, brr) == false {
			sort.Succeed = false
			break
		}
	}
	if sort.Succeed != true {
		log.Printf("merge bad")
		return
	}
	log.Printf("merge good")

	// 小数之和
	for i := 0; i < sort.TestTime; i++ {
		arr := sort.GenerateRandomArray(sort.Size, sort.Value)
		brr := sort.CopyArr(arr)
		aSum := smallSum.SmallSum(arr)
		bSum := smallSum.Comparator(brr)
		if aSum != bSum {
			sort.Succeed = false
			break
		}
	}
	if sort.Succeed != true {
		log.Printf("small sum bad")
		return
	}
	log.Printf("small sum good")
}
