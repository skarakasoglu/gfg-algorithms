package sort

// InsertionSort splits array into sorted and unsorted parts. Values from unsorted parts are placed the correct place in the sorted part.
// Time Complexity: O(N*N)
// Auxiliary Space: O(1)
// Stable: Yes
// In Place: Yes
// Algorithmic Paradigm: Incremental Approach
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i - 1] {
			for j := i - 1; j >= 0; j-- {
				if arr[j + 1] < arr[j] {
					arr[j + 1], arr[j] = arr[j], arr[j + 1]
				} else {
					break
				}
			}
		}
	}

	return arr
}
