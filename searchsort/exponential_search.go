package searchsort

// ExponentialSearch finds range where the element exists and does the binary search within that range
// works better than binary search when the element being searched is closer to the first element.
// Time Complexity: O(Log N)
func ExponentialSearch(arr []int, x int) int {
	i := 1

	// create subarrays with size of the powers of number 2.
	// when the last element of the array is greater or equal then the element being searched
	// then exit from the loop
	for ; i < len(arr) && arr[i] < x; i *= 2 {}

	if i >= len(arr) {
		i = len(arr) - 1
	}

	// we know that x is greater than elements before i / 2
	// and lower than the elements after i
	// do standard binary search
	low := (i / 2)
	high := i

	for high >= low {
		medium := (high + low) / 2

		if arr[medium] == x {
			return medium
		} else if arr[medium] > x {
			high = medium - 1
		} else {
			low = medium + 1
		}
	}

	return -1
}
