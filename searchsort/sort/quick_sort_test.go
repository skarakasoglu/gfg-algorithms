package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{4, 1, 3, 9, 7}

	QuickSort(arr)
	expected := []int{1,3,4,7,9}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Quick sort is incorrect, got: %d, expected: %v", arr, expected)
	}
}

func bechmarkQuickSort(arr []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		QuickSort(arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := []int{4, 1, 3, 9, 7}
	bechmarkQuickSort(arr, b)
}

func BenchmarkQuickSortWithLargeInputSorted(b *testing.B) {
	size := 1000000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}

	bechmarkQuickSort(arr, b)
}
