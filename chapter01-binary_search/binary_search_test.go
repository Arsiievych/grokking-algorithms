package chapter01_binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	got, _ := BinarySearch(arr, 12)
	want := 11

	if got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	num, ok := BinarySearch(arr, 21)

	if num != -1 {
		t.Errorf("expected -1, got %d", num)
	}

	if ok {
		t.Errorf("expected false, got true")
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 10000000)
	for i := range arr {
		arr[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = BinarySearch(arr, 9999999)
	}
}
