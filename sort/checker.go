package sort

import (
	"math/rand"
	"sort"
)

const (
	TestTime = 50000
	Size     = 1000
	Value    = 100
)

var (
	Succeed = true
)

func RightMethod(arr []int) {
	sort.Ints(arr)
}

func GenerateRandomArray(size, value int) []int {
	arr := make([]int, rand.Intn(size+1))
	for i := 0; i < len(arr); i++ {
		arr[i] = (value+1)*rand.Int() - value*rand.Int()
	}
	return arr
}

func CopyArr(arr []int) []int {
	brr := make([]int, len(arr))
	copy(brr, arr)
	return brr
}

func IsEqual(arr, brr []int) bool {
	if len(arr) != len(brr) {
		return false
	}
	for i, v := range arr {
		if v != brr[i] {
			return false
		}
	}
	return true
}
