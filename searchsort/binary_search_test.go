package searchsort

import "testing"

func TestBinarySearchWithAnExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23
	max := len(arr) - 1
	min := 0

	expectedIndex := 5

	index := BinarySearch(arr, x, max, min)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestBinarySearchWithNonExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 53
	max := len(arr) - 1
	min := 0

	expectedIndex := -1

	index := BinarySearch(arr, x, max, min)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkBinarySearch(arr []int, x int, max int, min int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		BinarySearch(arr, x, max, min)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23
	max := len(arr) - 1
	min := 0

	benchmarkBinarySearch(arr, x, max, min, b)
}

func TestBitwiseBinarySearchWithAnExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23

	expectedIndex := 5

	index := BitwiseBinarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestBitwiseBinarySearchWithNonExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 53

	expectedIndex := -1

	index := BitwiseBinarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestBitwiseBinarySearch(t *testing.T) {
	arr := []int{2,5,8,23}
	x := 23

	expectedIndex := 3

	index := BitwiseBinarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkBitwiseBinarySearch(arr []int, x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		BitwiseBinarySearch(arr, x)
	}
}

func BenchmarkBitwiseBinarySearch(b *testing.B) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23

	benchmarkBitwiseBinarySearch(arr, x, b)
}

