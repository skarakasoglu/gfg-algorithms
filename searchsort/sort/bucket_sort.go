package sort

import (
	"math"
)

// BucketSort
// Time Complexity: O(n)
func BucketSort(arr []float32) {
	n := len(arr) / 2

	buckets := make([][]float32, n)

	max := getMaxBucketSort(arr)
	min := getMinBucketSort(arr)

	elRange := int(math.Ceil(float64(max - min) / float64(n)))

	for i := range arr {
		bucketIndex := int((arr[i] - min) / float32(elRange))
		buckets[bucketIndex] = append(buckets[bucketIndex], arr[i])
	}

	for _, bucket := range buckets {
		for i := 1; i < len(bucket); i++ {
			if bucket[i] < bucket[i - 1] {
				for j := i - 1; j >= 0; j-- {
					if bucket[j + 1] < bucket[j] {
						bucket[j], bucket[j + 1] = bucket[j + 1], bucket[j]
					} else {
						break
					}
				}
			}
		}
	}

	currentIndex := 0
	for _, bucket := range buckets {
		for _, el := range bucket {
			arr[currentIndex] = el
			currentIndex++
		}
	}
}

func getMaxBucketSort(arr []float32) float32 {
	max := arr[0]

	for _, el := range arr {
		if max < el {
			max = el
		}
	}

	return max
}

func getMinBucketSort(arr []float32) float32 {
	min := arr[0]

	for _, el := range arr {
		if min > el {
			min = el
		}
	}

	return min
}