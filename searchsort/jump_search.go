package searchsort

import "math"

func JumpSearch(arr []int, x int) int {
	// block size to be jumped
	m := int(math.Sqrt(float64(len(arr))))

	i := 0
	max := 0

	// adding m to i in every iteration to jump
	for ; i < len(arr); i += m {

		// if current index has greater value than go back by m and do linear search to current index.
		// so store current index in max variable then break from the loop.
		if arr[i] > x {
			max = i
			i -= m
			break
		}
	}

	// Linear search to the max value
	for ; i < max; i++ {
		if arr[i] == x {
			return i
		}
	}

	return -1
}
