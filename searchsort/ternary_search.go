package searchsort

// TernarySearch it is a searching algorithm similar to binary search but it divides the array into 3 parts.
// Time complexity: O(Log3 N)
func TernarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for high >= low {
		medium1 := low + (high - low) / 3
		medium2 := high - (high - low) / 3

		if arr[medium1] == x {
			return medium1
		} else if arr[medium2] == x {
			return medium2
		}

		if arr[medium1] > x {
			high = medium1 - 1
		} else if arr[medium2] < x {
			low = medium2 + 1
		} else {
			low = medium1 + 1
			high = medium2 - 1
		}
	}

	return -1
}
