package main

import (
	"fmt"
	"os"
	"errors"
)

// Error wrapping allows you to add more context to an error without losing the original error.
// This creates a chain of errors that can be inspected later.

// A function that simulates reading a file.
func readFile(filename string) error {
	// os.ReadFile returns an error if the file doesn't exist.
	_, err := os.ReadFile(filename)
	if err != nil {
		// Use `fmt.Errorf` with `%w` to wrap the original error.
		return fmt.Errorf("failed to read file '%s': %w", filename, err)
	}
	return nil
}

// A function that uses the `readFile` function and may return its own error.
func processData() error {
	// Let's call the function with a non-existent file.
	err := readFile("non-existent.txt")
	if err != nil {
		return fmt.Errorf("could not process data: %w", err)
	}
	return nil
}

func main() {

	// --- Error Wrapping Example ---
	fmt.Println("--- Error Wrapping ---")
	err := processData()
	if err != nil {
		fmt.Println("Complete error chain:", err)

		// You can use `errors.Is` to check if an error in the chain matches a specific sentinel error.
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("\tDetected original error: file does not exist.")
		}

		// You can use `errors.Unwrap` to get the next error in the chain.
		unwrappedErr := errors.Unwrap(err)
		fmt.Println("\tUnwrapped error (1st level):", unwrappedErr)

		unwrappedErr = errors.Unwrap(unwrappedErr)
		fmt.Println("\tUnwrapped error (2nd level):", unwrappedErr)

	}
}
