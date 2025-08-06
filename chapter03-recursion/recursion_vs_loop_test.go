package chapter03_recursion

import "testing"

func TestSumLoop(t *testing.T) {
	want := 5050
	got := SumLoop(100)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func BenchmarkSumLoop(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SumLoop(10_000)
	}
}

func TestSumRecursion(t *testing.T) {
	want := 5050
	got := SumRecursion(100)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func BenchmarkSumRecursion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumRecursion(10_000)
	}
}

func TestGaussFormula(t *testing.T) {
	want := 5050
	got := GaussFormula(100)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func BenchmarkGaussFormula(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GaussFormula(10_000)
	}
}
