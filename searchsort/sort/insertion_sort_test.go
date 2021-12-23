package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6}

	InsertionSort(arr)
	expected := []int{5,6,11,12,13}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Insertion sort is incorrect, got: %d, expected: %v", arr, expected)
	}
}

func TestInsertionSort2(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	InsertionSort(arr)
	expected := []int{1,2,3,4,5,6,7,8,9,10}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Insertion sort is incorrect, got: %d, expected: %v", arr, expected)
	}
}

func benchmarkInsertionSort(arr []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		InsertionSort(arr)
	}
}
func BenchmarkInsertionSort(b *testing.B) {
	arr := []int{4, 1, 3, 9, 7}
	benchmarkMergeSort(arr, b)
}

func BenchmarkInsertionSortWithLargeInputSorted(b *testing.B) {
	size := 100000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}

	benchmarkInsertionSort(arr, b)
}