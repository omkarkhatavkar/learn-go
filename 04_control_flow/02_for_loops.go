package main

import "fmt"

func main() {

	// The `for` loop is the only loop construct in Go.
	// --- C-style `for` loop ---
	fmt.Println("--- C-style `for` loop ---")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// --- `for` as a `while` loop ---
	// You can omit the pre and post statements.
	fmt.Println("\n--- `for` as a `while` loop ---")
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	// --- Infinite `for` loop ---
	// You can omit all parts of the loop statement.
	// This is useful for event loops or daemons.
	fmt.Println("\n--- Infinite loop (Break with Ctrl+C) ---")
	counter := 0
	for {
		fmt.Println("Looping...")
		counter++
		if counter >= 3 {
			fmt.Println("Breaking out of the infinite loop.")
			break // Use `break` to exit a loop.
		}
	}

	// --- `for-range` loop ---
	// The `for-range` loop is used to iterate over elements of a slice, array, map, etc.
	fmt.Println("\n--- `for-range` loop ---")
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Use the blank identifier `_` to ignore a value you don't need.
	fmt.Println("\n--- `for-range` with `_` ---")
	for _, value := range slice {
		fmt.Println("Value:", value)
	}

	// `continue` jumps to the next iteration of the loop.
	fmt.Println("\n--- `continue` statement ---")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skip the rest of the loop for i=2
		}
		fmt.Println(i)
	}
}
