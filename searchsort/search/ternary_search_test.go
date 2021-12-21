package search

import "testing"

func TestTernarySearch(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23

	expectedIndex := 5

	index := TernarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestTernarySearchWithLargeInput(t *testing.T) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 5731

	expectedIndex := 5730

	index := TernarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkTernarySearch(arr []int, x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		TernarySearch(arr, x)
	}
}

func BenchmarkTernarySearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	benchmarkTernarySearch(arr, x, b)
}