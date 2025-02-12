package main

import "fmt"

func main() {

	// --- Basic `if` statement ---
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5.")
	}

	// --- `if` with an `else` clause ---
	score := 85
	if score >= 90 {
		fmt.Println("You got an A.")
	} else {
		fmt.Println("You did not get an A.")
	}

	// --- `if/else if/else` chain ---
	grade := 75
	if grade >= 90 {
		fmt.Println("Grade: A")
	} else if grade >= 80 {
		fmt.Println("Grade: B")
	} else if grade >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: Below C")
	}

	// --- `if` with a short statement ---
	// This is a common Go idiom. You can declare and initialize a variable
	// within the `if` statement's scope.
	// The variable `result` is only available within the `if` and `else` blocks.
	fmt.Println("\n--- `if` with short statement ---")
	// `len` returns the length of a string
	if result := len("Hello, World!"); result > 10 {
		fmt.Printf("The string is long. Length: %d\n", result)
	} else {
		fmt.Printf("The string is short. Length: %d\n", result)
	}

	// You can't access `result` here because it's out of scope.
	// fmt.Println(result) // This would cause a compilation error.
}
