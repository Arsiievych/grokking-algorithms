package chapter04_quick_sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "empty array",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "one element array",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "two elements in order",
			arr:  []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "two elements out of order",
			arr:  []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "array with duplicates",
			arr:  []int{8, 2, 3, 5, 1, 3, 4, 6, 7, 9, 10, 2, 4, 3},
			want: []int{1, 2, 2, 3, 3, 3, 4, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "already sorted array",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse sorted array",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := QuickSort(tc.arr)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tc.want)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = 10000 - i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr)
	}
}
