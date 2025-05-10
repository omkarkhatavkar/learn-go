package main

import "testing"

// To run the tests, use the command:
// go test
// To see verbose output, use:
// go test -v

// A basic unit test function. The function name must start with `Test` and take one argument of type `*testing.T`.
func TestSum(t *testing.T) {

	// Test case 1: A simple positive case.
	numbers := []int{1, 2, 3}
	expected := 6
	actual := Sum(numbers)

	if actual != expected {
		// t.Errorf is used to report a test failure.
		t.Errorf("Test failed: Expected %d, but got %d", expected, actual)
	}

	// Test case 2: Empty slice.
	numbers = []int{}
	expected = 0
	actual = Sum(numbers)

	if actual != expected {
		t.Errorf("Test failed for empty slice: Expected %d, but got %d", expected, actual)
	}

	// Test case 3: Negative numbers.
	numbers = []int{-1, -2, -3}
	expected = -6
	actual = Sum(numbers)

	if actual != expected {
		t.Errorf("Test failed for negative numbers: Expected %d, but got %d", expected, actual)
	}
}
