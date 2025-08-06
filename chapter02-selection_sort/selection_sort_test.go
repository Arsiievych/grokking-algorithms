package chapter02_selection_sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{8, 2, 3, 5, 1, 3, 4, 6, 7, 9, 10, 2, 4, 3}

	want := []int{1, 2, 2, 3, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10}
	got := SelectionSort(arr)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SelectionSort() = %v, want %v", got, want)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = 10000 - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}

func TestFindMin(t *testing.T) {
	t.Run("normal array", func(t *testing.T) {
		arr := []int{64, 25, 12, 22, 11}
		want := 4
		got := findMin(arr)
		if got != want {
			t.Errorf("findMin() = %v, want %v", got, want)
		}
	})

	t.Run("empty array", func(t *testing.T) {
		arr := []int{}
		want := -1
		got := findMin(arr)
		if got != want {
			t.Errorf("findMin() = %v, want %v", got, want)
		}
	})

	t.Run("array with duplicate min values", func(t *testing.T) {
		arr := []int{64, 25, 12, 11, 11}
		want := 3
		got := findMin(arr)
		if got != want {
			t.Errorf("findMin() = %v, want %v", got, want)
		}
	})

	t.Run("single element array", func(t *testing.T) {
		arr := []int{5}
		want := 0
		got := findMin(arr)
		if got != want {
			t.Errorf("findMin() = %v, want %v", got, want)
		}
	})
}
