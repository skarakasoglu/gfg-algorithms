package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 1, 3, 9, 7}

	got := BubbleSort(arr)
	expected := []int{1,3,4,7,9}

	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
		t.Errorf("Bubble sort is incorrect, got: %d, expected: %v", got, expected)
	}
}

func TestBubbleSort2(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	got := BubbleSort(arr)
	expected := []int{1,2,3,4,5,6,7,8,9,10}

	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
		t.Errorf("Bubble sort is incorrect, got: %d, expected: %v", got, expected)
	}
}

func benchmarkBubbleSort(arr []int, b *testing.B) {
	for n := 0; n < b.N; n++{
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	benchmarkBubbleSort(arr, b)
}