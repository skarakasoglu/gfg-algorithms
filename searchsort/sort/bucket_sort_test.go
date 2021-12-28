package sort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []float32{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
	BucketSort(arr)

	expected := []float32{0.12, 0.17, 0.21, 0.23, 0.26, 0.39, 0.68, 0.72, 0.78, 0.94}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Bucket sort is incorrect, got: %v, expected: %v", arr, expected)
	}
}

func TestBucketSort2(t *testing.T) {
	arr := []float32{9.8 , 0.6 , 10.1 , 1.9 , 3.07 , 3.04 , 5.0 , 8.0 , 4.8 , 7.68}
	BucketSort(arr)

	expected := []float32{0.6 , 1.9 , 3.04 , 3.07 , 4.8 , 5.0 , 7.68 , 8.0 , 9.8 , 10.1}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Bucket sort is incorrect, got: %v, expected: %v", arr, expected)
	}
}

func TestBucketSort3(t *testing.T) {
	arr := []float32{0.49 , 5.9 , 3.4 , 1.11 , 4.5 , 6.6 , 2.0}
	BucketSort(arr)

	expected := []float32{0.49 , 1.11 , 2.0 , 3.4 , 4.5 , 5.9 , 6.6}

	if fmt.Sprintf("%v", arr) != fmt.Sprintf("%v", expected) {
		t.Errorf("Bucket sort is incorrect, got: %v, expected: %v", arr, expected)
	}
}