package sort

import (
	"fmt"
	"testing"
)

func TestRadixSortUsingCountingSort(t *testing.T) {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
    RadixSort(arr, RadixSortType_CountingSort)

	expected := []int{2, 24, 45, 66, 75, 90, 170, 802}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Bucket sort test failed. Expected: %v, got: %v", expected, arr)
	}
}
