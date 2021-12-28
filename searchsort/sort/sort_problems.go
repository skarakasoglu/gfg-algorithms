package sort

// LengthUnsortedSubarray Given an unsorted array arr[0..n-1] of size n,
// find the minimum length subarray arr[s..e] such that sorting this subarray makes the whole array sorted.
// returns left and right indexes
func LengthUnsortedSubarray(arr []int) (int, int) {
	left := 0
	right := len(arr) - 1

	leftIndex := -1
	rightIndex := -1

	// Finding candidate unsorted subarray
	for left < right {
		if arr[left] < arr[left + 1] {
			left++
		} else {
			leftIndex = left
		}

		if arr[right - 1] < arr[right] {
			right--
		} else {
			rightIndex = right
		}

		if leftIndex != -1 && rightIndex != -1 {
			break
		}
	}

	min := arr[leftIndex]
	max := arr[rightIndex]

	// Finding minimum ve maximum values of the subarray
	for i := leftIndex; i <= rightIndex; i++ {
		if min > arr[i] {
			min = arr[i]
		}

		if max < arr[i] {
			max = arr[i]
		}
	}

	// Finding greater values than minimum values of the subarray in starting to left, if so the subarray will start from this index.
	for i := 0; i < leftIndex; i++ {
		if arr[i] > min {
			leftIndex = i
			break
		}
	}

	// Finding lower values than maximum values of the subarray in starting from right to end, if so the subarray will end at this index.
	for i := rightIndex + 1; i < len(arr); i++ {
		if arr[i] < max {
			rightIndex = i
			break
		}
	}

	return leftIndex, rightIndex
}