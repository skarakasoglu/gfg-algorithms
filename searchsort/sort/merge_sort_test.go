package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{4, 1, 3, 9, 7}

	got := MergeSort(arr)
	expected := []int{1,3,4,7,9}

	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
		t.Errorf("Merge sort is incorrect, got: %d, expected: %v", got, expected)
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

func BenchmarkMergeSortWithLargeInput(b *testing.B) {
	size := 100000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}

	benchmarkMergeSort(arr, b)
}