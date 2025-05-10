package main

// This is a simple function that we will write tests for.
// The test file for this will be named `sum_test.go`.

// Sums a slice of integers.
func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
