package sort

import "testing"

func TestLengthUnsortedSubarray(t *testing.T) {
	arr := []int{10, 12, 20, 30, 25, 40, 32, 31, 35, 50, 60}
	left, right := LengthUnsortedSubarray(arr)

	expectedLeft, expectedRight := 3, 8

	if left != expectedLeft || right != expectedRight {
		t.Errorf("Finding unsorted subarray failed. Expected: %v %v, got: %v %v", expectedLeft, expectedRight, left, right)
	}
}

func TestLengthUnsortedSubarray2(t *testing.T) {
	arr := []int{0, 1, 15, 25, 6, 7, 30, 40, 50}
	left, right := LengthUnsortedSubarray(arr)

	expectedLeft, expectedRight := 2, 5

	if left != expectedLeft || right != expectedRight {
		t.Errorf("Finding unsorted subarray failed. Expected: %v %v, got: %v %v", expectedLeft, expectedRight, left, right)
	}
}
