package search

import "testing"

func TestExponentialSearch(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23

	expectedIndex := 5

	index := ExponentialSearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestExponentialSearchWithLargeInput(t *testing.T) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 8765

	expectedIndex := 8764

	index := ExponentialSearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkExponentialSearch(arr []int, x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		ExponentialSearch(arr, x)
	}
}

func BenchmarkExponentialSearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	benchmarkExponentialSearch(arr, x, b)
}