package sort

import (
	"fmt"
	"testing"
)

func TestKSortedArray(t *testing.T) {
	arr := []int{6, 5, 3, 2, 8, 10, 9}
	k := 3

	KSortedArray(arr, k)

	expected := []int{2, 3, 5, 6, 8, 9, 10}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Sorting K sorted array is incorrect, got: %d, expected: %v", arr, expected)
	}
}

func TestKSortedArray2(t *testing.T) {
	arr := []int{10, 9, 8, 7, 4, 70, 60, 50}
	k := 4

	KSortedArray(arr, k)

	expected := []int{4, 7, 8, 9, 10, 50, 60, 70}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Sorting K sorted array is incorrect, got: %d, expected: %v", arr, expected)
	}
}