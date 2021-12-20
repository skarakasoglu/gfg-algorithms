package searchsort

import "math"

// JumpSearch jumps ahead at fixed steps if the current index value is greater than x
// then goes back for one fixed step and does linear search.
// Time complexity: O(√n)
func JumpSearch(arr []int, x int) int {
	// block size to be jumped
	// n / m: it skips m values in every turn, m - 1: in the worst case it does linear search m - 1 times.
	// in the worst case the total number of comparisons will be ((n / m) + (m - 1))
	// to be able to find the best m value, this equation should be minimum.
	// the minimum and maximum values are where the derivative of an equation change sign and equal to 0.
	// if we take the derivative of this equation and find the root, it will be m = √n
	// therefore the best block size is √n
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
