package main

import (
	"errors"
	"fmt"
)

// Go functions can return multiple values. This is commonly used for returning a
// result and an error value simultaneously.
// A function to divide two numbers, returning the result and a potential error.
func divide(x, y float64) (float64, error) {
	// Check for division by zero.
	if y == 0 {
		// The `errors.New` function creates a new error with a given message.
		return 0, errors.New("cannot divide by zero")
	}

	// If no error occurred, return the result and a nil error.
	// `nil` is Go's way of representing a null value, often used for errors.
	return x / y, nil
}

// You can also name the return values. This makes the function clearer and
// allows for a "naked" return, which simply returns the named values.
func namedReturn(x, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y
	// A naked return returns the current values of `sum` and `diff`.
	return
}

func main() {

	// --- Example with division ---
	// Call the function and handle both return values.
	result1, err1 := divide(10, 2)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Printf("10 / 2 = %f\n", result1)
	}

	// Example of an error case.
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("\nError:", err2)
	} else {
		fmt.Printf("10 / 0 = %f\n", result2)
	}

	// --- Example with named returns ---
	fmt.Println("\n--- Named Returns ---")
	sum, diff := namedReturn(8, 2)
	fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)
}
