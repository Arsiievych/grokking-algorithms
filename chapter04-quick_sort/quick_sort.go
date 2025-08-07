package chapter04_quick_sort

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less, greater := groupByPivot(arr, pivot)

	result := append(QuickSort(less), pivot)
	result = append(result, QuickSort(greater)...)
	return result
}

func groupByPivot(arr []int, pivot int) (less []int, greater []int) {
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return
}
