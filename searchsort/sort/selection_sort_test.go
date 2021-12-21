package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 1, 3, 9, 7}

	got := SelectionSort(arr)
	expected := []int{1,3,4,7,9}

	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
		t.Errorf("Selection sort is incorrect, got: %d, expected: %v", got, expected)
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