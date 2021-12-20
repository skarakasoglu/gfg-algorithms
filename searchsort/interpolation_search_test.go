package searchsort

import "testing"

func TestInterpolationSearchWithAnExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 23

	expectedIndex := 5

	index := InterpolationSearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}

func TestInterpolationSearchWithNonExistingElement(t *testing.T) {
	arr := []int{2,5,8,12,16,23,38,56,72,91}
	x := 89

	expectedIndex := -1

	index := InterpolationSearch(arr, x)

	if index != expectedIndex {
		t.Errorf("Element index is incorrect, got: %d, expected: %v", index, expectedIndex)
	}
}


func benchmarkInterpolationSearch(arr []int, x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		InterpolationSearch(arr, x)
	}
}

func BenchmarkInterpolationSearch(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 1; i <= size; i++ {
		arr[i - 1] = i
	}
	x := 9999

	benchmarkInterpolationSearch(arr, x, b)
}