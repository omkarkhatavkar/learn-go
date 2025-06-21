package main

import (
	"errors"
	"fmt"
)

// Go's idiomatic approach to error handling is to return an error as the last return value.

// A simple function that returns a result and an error.
func getNumber(input int) (int, error) {
	// If the input is valid, return the number and a nil error.
	if input > 0 {
		return input, nil
	}

	// If the input is invalid, return the zero value and a new error.
	// `errors.New` creates a simple error message.
	return 0, errors.New("input must be a positive number")
}

// You can also create custom error types by defining a new struct that implements the `Error()` method.
// This is useful for providing more context or structured error data.

// Define a custom error struct.
type customError struct {
	Message string
	Code    int
}

// Implement the `Error()` method to satisfy the `error` interface.
func (e *customError) Error() string {
	return fmt.Sprintf("custom error: code %d, message: %s", e.Code, e.Message)
}

func performOperation(value int) error {
	if value < 0 {
		return &customError{
			Message: "value cannot be negative",
			Code:    400,
		}
	}
	fmt.Println("Operation successful!")
	return nil
}

func main() {
	// --- Simple Error Example ---
	fmt.Println("--- Simple Error Handling ---")
	val1, err := getNumber(10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Received number:", val1)
	}

	val2, err := getNumber(-5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Received number:", val2)
	}

	// --- Custom Error Example ---
	fmt.Println("\n--- Custom Error Handling ---")
	err = performOperation(-10)
	if err != nil {
		fmt.Println("Caught custom error:", err)
	
		// You can use a type assertion to check if the error is of your custom type.
		var ce *customError
		// `errors.As` is the preferred way to unwrap and check for a specific error type.
		if errors.As(err, &ce) {
			fmt.Printf("\tThis is a custom error! Code: %d, Message: %s\n", ce.Code, ce.Message)
		}
	}
	
	err = performOperation(10)
	if err != nil {
		fmt.Println("Caught custom error:", err)
	}
}
