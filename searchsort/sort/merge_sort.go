package sort

// MergeSort divides the array into two halves recursively until the size of a halve becomes 1.
// Then merge halves till the complete array is merged.
// Time Complexity: O(n Log n)
// Auxiliary Space: O(n)
// Algorithm Paradigm: Divide and Conquer
// Stable: Yes
// In Place: No
func MergeSort(arr []int) {
	l := 0
	r := len(arr) - 1

	mergeSort(arr, l, r)
}

func mergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	medium := left + (right - left) / 2

	mergeSort(arr, left, medium)
	mergeSort(arr, medium + 1, right)
	merge(arr, left, medium, right)
}

func merge(arr []int, left int, medium int, right int) {
	leftArraySize := medium - left + 1
	rightArraySize := right - medium

	leftArray := make([]int, leftArraySize)
	rightArray := make([]int, rightArraySize)

	_ = copy(leftArray, arr[left: left + leftArraySize])
	_ = copy(rightArray, arr[medium + 1: medium + 1 + rightArraySize])

	leftArrayIndex := 0
	rightArrayIndex := 0
	mergedArrayIndex := left

	for  rightArrayIndex < rightArraySize && leftArrayIndex < leftArraySize {
		if leftArray[leftArrayIndex] <= rightArray[rightArrayIndex] {
			arr[mergedArrayIndex] = leftArray[leftArrayIndex]
			leftArrayIndex++
		} else {
			arr[mergedArrayIndex] = rightArray[rightArrayIndex]
			rightArrayIndex++
		}
		mergedArrayIndex++
	}

	for leftArrayIndex < leftArraySize {
		arr[mergedArrayIndex] = leftArray[leftArrayIndex]
		leftArrayIndex++
		mergedArrayIndex++
	}

	for rightArrayIndex < rightArraySize {
		arr[mergedArrayIndex] = rightArray[rightArrayIndex]
		rightArrayIndex++
		mergedArrayIndex++
	}
}