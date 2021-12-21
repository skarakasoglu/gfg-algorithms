package sort

// SelectionSort sort an array by iterating through and placing the minimum element to the start
// Time Complexity: O(N2)
// Auxiliary Space: O(1)
// Stable: Default implementation is not stable, but it can be made.
// In place: Yes, does not require extra space.
func SelectionSort(arr []int) []int {

	subIndex := 0

	for subIndex < len(arr) {
		currentMinimum := arr[subIndex]
		currentMinimumIndex := subIndex

		for i := subIndex; i < len(arr); i++{
			if arr[i] < currentMinimum {
				currentMinimum = arr[i]
				currentMinimumIndex = i
			}
		}

		arr[subIndex], arr[currentMinimumIndex] = arr[currentMinimumIndex], arr[subIndex]
		subIndex++
	}

	return arr
}
