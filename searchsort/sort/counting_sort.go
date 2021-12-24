package sort

import "strings"

// LargestNumber Given an integer S. Finds the largest even number that can be formed by rearranging the digits of S.
// In case the number does not contain any even digit then outputs the largest odd number possible.
func LargestNumber(s string) string {
	// Converting string characters to integer array so that we can give it as an argument to CountingSort method.
	arr := make([]int, len(s))
	for i, ch := range s {
		arr[i] = int(ch - '0')
	}
	outputInt := CountingSort(arr)

	// Converting sorted output array to chars again to use it in string
	output := make([]rune, len(outputInt))
	for i, digit := range outputInt {
		output[i] = rune(digit + '0')
	}

	// Finding the the smallest even number in the string.
	evenNumberIndex := -1
	for i, val := range output {
		digit := int(val - '0')

		// When we found the first even digit in the array, we should exit from the loop,
		// because we need to find the largest even number.
		// even if there are greater even numbers, these should be used in greater digits like 10s 100s etc.
		if digit % 2 == 0{
			evenNumberIndex = i
			break
		}
	}

	largestEven := make([]rune, len(output))

	// If there is an even number, we should run the loop one time less.
	// because we will have been already placed a digit to our number.
	max := len(output)
	if evenNumberIndex != -1 {
		largestEven[len(largestEven) - 1] = output[evenNumberIndex]
		max--
	}

	currentIndex := len(output) - 1
	for i := 0; i < max; i++ {

		if currentIndex == evenNumberIndex {
			currentIndex--
			i--
			continue
		}

		largestEven[i] = output[currentIndex]
		currentIndex--
	}

	var x strings.Builder
	for _, val := range largestEven {
		x.WriteRune(val)
	}

	return x.String()
}

// CountingSort counts the number of objects having distinct key values, calculates the position of each object in the output sequence.
// Time Complexity: O(n + k)
// Auxiliary Space: O(n + k)
func CountingSort(arr []int) []int {
	frequencies := make([]int, 9)
	for i, _ := range frequencies {
		frequencies[i] = 0
	}

	for _, digit := range arr {
		frequencies[digit]++
	}

	for i := 1; i < len(frequencies); i++ {
		frequencies[i] += frequencies[i - 1]
	}

	output := make([]int, len(arr))

	for _, digit := range arr {
		output[frequencies[digit] - 1] = digit
		frequencies[digit]--
	}

	return output
}
