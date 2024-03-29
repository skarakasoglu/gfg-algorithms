package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 1, 3, 9, 7}

	SelectionSort(arr)
	expected := []int{1,3,4,7,9}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Selection sort is incorrect, got: %d, expected: %v", arr, expected)
	}
}

func benchmarkSelectionSort(arr []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelectionSort(arr)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	arr := []int{4, 1, 3, 9, 7}
	benchmarkSelectionSort(arr, b)
}

func BenchmarkSelectionSortWithLargeInputSorted(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}

	benchmarkSelectionSort(arr, b)
}