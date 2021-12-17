package searchsort

// Question: Given a sorted array arr[] of n elements, write a function to search a given element x in arr[].

// BinarySearch Recursive binary search function to find the element x in given array arr.
// arr: the array we look for element x in
// x: the element we search
// max: the maximum index of the elements that have possibility to equal to x
// min: the minimum index of the elements that have possibility to equal to x
func BinarySearch(arr []int, x int, max int, min int) int {
	medium := (max + min) / 2

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
