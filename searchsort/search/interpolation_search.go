package search

// InterpolationSearch it is a binary search logically, but binary search takes medium value as checkpoints
// this method goes to different locations to the position of the element being searched.
// for ex. if the element is close to the end of the array then the search will likely to be started toward the end of the array.
// it is better when the values are uniformly distributed.
// the position is calculated using the probe position formula from geometry.
// Average case: O(Log Log N), Worst Case: O(N)
func InterpolationSearch(arr []int, x int) int {
	high := len(arr) - 1
	low := 0

	for high >= low {
		// If low equals to high in the position formula the value will result to low
		// Therefore, in this case no need to calculate, if they are equal
		// Then, check if the index value equals to x, if so return index otherwise return -1.
		if low == high {
			if arr[low] == x {
				return low
			}

			return -1
		}

		// calculating the probe position formula
		pos := low + (x - arr[low]) * (high - low) / (arr[high] - arr[low])

		if arr[pos] == x {
			return pos
		}

		if arr[pos] > x {
			high = pos - 1
		} else {
			low = pos + 1
		}
	}

	return -1
}