package sort

// HeapSort sort the array using binary heap data structure. (complete binary tree)
// Time Complexity: O(n Log n)
// Auxiliary Space: O(1)
// Stable: No
// In place: Yes
func HeapSort(arr []int) []int {
	return buildHeap(arr)
}

func buildHeap(arr []int) []int {
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		heapify(arr, i, 0)
	}

	return arr
}

func heapify(arr []int, n int, root int) []int{
	left := 2 * root + 1
	right := 2 * root + 2

	largest := root

	if left < n && arr[largest] < arr[left] {
		largest = left
	}

	if right < n && arr[largest] < arr[right] {
		largest = right
	}

	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, n, largest)
	}

	return arr
}