package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{4, 1, 3, 9, 7}

	MergeSort(arr)
	expected := []int{1,3,4,7,9}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Merge sort is incorrect, got: %d, expected: %v", arr, expected)
	}
}

func benchmarkMergeSort(arr []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	arr := []int{4, 1, 3, 9, 7}
	benchmarkMergeSort(arr, b)
}

func BenchmarkMergeSortWithLargeInputSorted(b *testing.B) {
	size := 1000000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}

	benchmarkMergeSort(arr, b)
}