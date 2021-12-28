package sort

// KSortedArray sorting a nearly (K sorted) sorted array with insertion sort.
func KSortedArray(arr []int, k int) {
	for i := 0; i < len(arr); i++ {
		for j := i - 1; j >= i - k && j >= 0; j-- {
			if arr[j + 1] < arr[j] {
				arr[j + 1], arr[j] = arr[j], arr[j + 1]
			} else {
				break
			}
		}

		for j := i + 1; j <= i + k && j < len(arr) - 1; j++ {
			if arr[j - 1] > arr[j] {
				arr[j + 1], arr[j] = arr[j], arr[j + 1]
			} else {
				break
			}
		}
	}
}
