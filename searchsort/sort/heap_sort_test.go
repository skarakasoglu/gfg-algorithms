package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{4, 1, 3, 7, 9}

	got := HeapSort(arr)
	expected := []int{1, 3, 4, 7, 9}

	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
		t.Errorf("Heap sort is incorrect, got: %d, expected: %v", got, expected)
	}
}

func benchmarkHeapSort(arr []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		HeapSort(arr)
	}
}

func BenchmarkHeapSortWithLargeInputSorted(b *testing.B) {
	size := 100000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}

	benchmarkHeapSort(arr, b)
}