package searchsort

import "testing"

func TestJumpSearchWithAnExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23

	expectedIndex := 5

	index := JumpSearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestJumpSearchWithNonExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 53

	expectedIndex := -1

	index := JumpSearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkJumpSearch(arr []int, x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		JumpSearch(arr, x)
	}
}

func BenchmarkJumpSearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i < size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	benchmarkJumpSearch(arr, x, b)
}