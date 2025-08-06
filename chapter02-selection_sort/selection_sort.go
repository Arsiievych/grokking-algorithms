package chapter02_selection_sort

func SelectionSort(arr []int) []int {
	sortedArr := make([]int, len(arr))

	for i := 0; i < len(sortedArr); i++ {
		minIdx := findMin(arr)
		sortedArr[i] = arr[minIdx]
		arr = append(arr[:minIdx], arr[minIdx+1:]...)
	}

	return sortedArr
}

func findMin(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	minIdx := 0
	minVal := arr[0]
	for idx, val := range arr {
		if val < minVal {
			minIdx = idx
			minVal = val
		}
	}

	return minIdx
}
