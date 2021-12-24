package sort

import (
	"testing"
)

func TestLargestNumberWithPossibleEvenNumber(t *testing.T) {
	num := LargestNumber("1453")
	expected := "5314"

	if expected != num {
		t.Errorf("Result is incorrect, got: %v, expected: %v", num, expected)
	}
}

func TestLargestNumberWithoutPossibleEvenNumber(t *testing.T) {
	num := LargestNumber("3555")
	expected := "5553"

	if expected != num {
		t.Errorf("Result is incorrect, got: %v, expected: %v", num, expected)
	}
}
