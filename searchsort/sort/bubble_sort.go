package sort

// BubbleSort swaps adjacent elements if the order is wrong
// Worst and Average Case Time Complexity: O(n*n)
// Best case Time Complexity: O(n)
// Auxiliary Space: O(1)
// In Place: Yes
// Stable: Yes
func BubbleSort(arr []int){
	for i := 0; i < len(arr) - 1; i++ {
		anySwap := false

		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				anySwap = true
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}

		if !anySwap {
			break
		}
	}
}
