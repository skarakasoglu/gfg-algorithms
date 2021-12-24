package sort

import (
	"math"
)

type RadixSortType int8

const (
	RadixSortType_CountingSort RadixSortType = iota
	RadixSortType_BucketSort
)

// RadixSort sorts digit by digit using counting sort or bucket sort.
// Time Complexity: O(logb(k)*(n+b)), b: numbers, k: maximum possible number
func RadixSort(arr []int, sortType RadixSortType) {
	max := getMax(arr)

	base := 10

	if sortType == RadixSortType_CountingSort {
		for i := 0; max / int(math.Pow(float64(base), float64(i))) > 0; i++ {
			countingSort(arr, base, i)
		}
	}
}

func countingSort(arr []int, base int, currentDigit int) {
	digitDivider := int(math.Pow(float64(base), float64(currentDigit)))
	frequencies := make([]int, 10)
	sorted := make([]int, len(arr))

	for _, number := range arr {
		digit := (number / digitDivider) % int(base)
		frequencies[digit]++
	}

	for j := 1; j < len(frequencies); j++ {
		frequencies[j] += frequencies[j - 1]
	}

	for j := len(arr) - 1; j >= 0; j--{
		digit := (arr[j] / digitDivider) % int(base)

		frequencies[digit]--
		sorted[frequencies[digit]] = arr[j]
	}

	for j := range sorted {
		arr[j] = sorted[j]
	}
}

func getMax(arr []int) int {

	max := arr[0]
	for _, val := range arr {
		if max < val {
			max = val
		}
	}

	return max
}