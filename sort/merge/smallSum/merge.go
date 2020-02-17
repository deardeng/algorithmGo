package smallSum

func MergeSort(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return sortProcess(arr, 0, len(arr)-1)
}

func sortProcess(arr []int, left, right int) int {
	if left == right {
		return 0
	}

	mid := left + (right-left)>>1
	return sortProcess(arr, left, mid) + sortProcess(arr, mid+1, right) + merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) int {
	helper := make([]int, right-left+1)
	i := 0
	p1, p2 := left, mid+1
	var result int
	for p1 <= mid && p2 <= right {
		if arr[p1] < arr[p2] {
			// 左边比右边小，左边arr[p1]将在小数之和中出现（right - p2 + 1）次
			result += (right - p2 + 1) * arr[p1]
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
	return result
}
