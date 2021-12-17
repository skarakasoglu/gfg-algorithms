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