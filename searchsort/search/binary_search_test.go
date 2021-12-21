package search

import "testing"

func TestBinarySearchWithAnExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23
	max := len(arr) - 1
	min := 0

	expectedIndex := 5

	index := BinarySearch(arr, x, max, min)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestBinarySearchWithNonExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 53
	max := len(arr) - 1
	min := 0

	expectedIndex := -1

	index := BinarySearch(arr, x, max, min)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestBinarySearchWithLargeInput(t *testing.T) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 4542

	max := len(arr) - 1
	min := 0

	expectedIndex := 4541

	index := BinarySearch(arr, x, max, min)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkBinarySearch(arr []int, x int, max int, min int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		BinarySearch(arr, x, max, min)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	max := len(arr) - 1
	min := 0

	benchmarkBinarySearch(arr, x, max, min, b)
}

func TestBitwiseBinarySearchWithAnExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23

	expectedIndex := 5

	index := BitwiseBinarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestBitwiseBinarySearchWithNonExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 53

	expectedIndex := -1

	index := BitwiseBinarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestBitwiseBinarySearch(t *testing.T) {
	arr := []int{2,5,8,23}
	x := 23

	expectedIndex := 3

	index := BitwiseBinarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkBitwiseBinarySearch(arr []int, x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		BitwiseBinarySearch(arr, x)
	}
}

func BenchmarkBitwiseBinarySearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	benchmarkBitwiseBinarySearch(arr, x, b)
}

func TestNonRecursiveBinarySearch(t *testing.T) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	expectedIndex := 9998
	index := NonRecursiveBinarySearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func benchmarkNonRecursiveBinarySearch(arr []int, x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		NonRecursiveBinarySearch(arr, x)
	}
}

func BenchmarkNonRecursiveBinarySearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	benchmarkNonRecursiveBinarySearch(arr, x, b)
}
