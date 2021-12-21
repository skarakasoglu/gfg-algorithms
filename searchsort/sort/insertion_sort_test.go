package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6}

	got := InsertionSort(arr)
	expected := []int{5,6,11,12,13}

	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
		t.Errorf("Insertion sort is incorrect, got: %d, expected: %v", got, expected)
	}
}

func TestInsertionSort2(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	got := InsertionSort(arr)
	expected := []int{1,2,3,4,5,6,7,8,9,10}

	if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected) {
		t.Errorf("Insertion sort is incorrect, got: %d, expected: %v", got, expected)
	}
}