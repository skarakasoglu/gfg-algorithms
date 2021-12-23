package sort

// QuickSort picks an element as pivot and partitions the array around the picked pivot.
// Time Complexity: Worst Case O(n^2), Average Case O(n Log n), Best Case O(n Log n)
// Stable: No
// In place: Yes
// Algoithm Paradigm: Divide and Conquer
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr) - 1)
}

func quickSort(arr []int, low int, high int) {
	if low < high {

		pi := partition(arr, low, high)

		quickSort(arr, low, pi - 1)
		quickSort(arr, pi + 1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high - 1; j++ {

		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i + 1], arr[high] = arr[high], arr[i + 1]

	return i + 1
}