package main

import "testing"

// Table-driven tests are a common Go idiom for writing robust and readable tests.
// It involves defining a slice of structs, where each struct represents a single test case.

// The `Sum` function is from the `sum.go` file in this package.

func TestTableDrivenSum(t *testing.T) {
	// A slice of structs, where each struct is a test case.
	// The fields are a descriptive name, the input, and the expected output.
	var tests = []struct {
		name     string
		input    []int
		expected int
	}{
		{"positive numbers", []int{1, 2, 3}, 6},
		{"empty slice", []int{}, 0},
		{"single number", []int{10}, 10},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed numbers", []int{5, -5, 10}, 10},
	}

	// Iterate over the slice of test cases.
	for _, tt := range tests {
		// Use `t.Run()` to run a subtest for each case.
		// This provides better isolation and more specific output in case of failure.
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test.
			actual := Sum(tt.input)

			// Check if the result matches the expected value.
			if actual != tt.expected {
				// Use `t.Errorf` to report a failure.
				t.Errorf("input %v: expected %d, got %d", tt.input, tt.expected, actual)
			}
		})
	}
}
