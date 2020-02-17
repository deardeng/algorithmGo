package smallSum

func SmallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return MergeSort(arr)
}

// 对数器，数组index i左边的数 比brr[i]小的数相加之和
func Comparator(brr []int) int {
	if brr == nil || len(brr) < 2 {
		return 0
	}

	var res int
	for i := 1; i < len(brr); i++ {
		for j := 0; j < i; j++ {
			if brr[j] < brr[i] {
				res += brr[j]
			}
		}
	}

	return res
}
