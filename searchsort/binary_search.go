package searchsort

// Question: Given a sorted array arr[] of n elements, write a function to search a given element x in arr[].

// BinarySearch Recursive binary search function to find the element x in given array arr.
// arr: the array we look for element x in
// x: the element we search
// max: the maximum index of the elements that have possibility to equal to x
// min: the minimum index of the elements that have possibility to equal to x
func BinarySearch(arr []int, x int, max int, min int) int {
	medium := min + (max - min) / 2

	if max >= min {
		if arr[medium] > x {
			return BinarySearch(arr, x, medium - 1, min)
		} else if arr[medium] < x {
			return BinarySearch(arr, x, max, medium + 1)
		} else if arr[medium] == x {
			return medium
		}
	}

	return -1
}

// NonRecursiveBinarySearch just a regular binary search implementation using loops.
func NonRecursiveBinarySearch(arr []int, x int) int {
	min := 0
	max := len(arr) - 1

	for max > min {
		medium := (min + max) / 2

		if arr[medium] > x {
			max = medium - 1
		} else if arr[medium] < x {
			min = medium + 1
		} else if arr[medium] == x {
			return medium
		}
	}

	return -1
}

// BitwiseBinarySearch Since all numbers can be represented as a sum of the powers of 2
// this function performs binary search using bitwise operators.
// arr: the array to look for in
// x: the element to be searched
func BitwiseBinarySearch(arr []int, x int) int {
	index := 0
	power := 1

	// Taking the power of the up to array size by taking the powers of 2, we will be taking the medium of the interval we should search in the upcoming section.
	for  ; power <= len(arr); power <<= 1 {}

	for power > 0 {
		// Diving power by 2 every time.
		power >>= 1

		// If index + power is smaller than x value then our minimum value should be greater.
		if index + power < len(arr) && arr[index + power] <= x {
			index += power
		}
	}

	if arr[index] == x {
		return index
	}

	return -1
}
