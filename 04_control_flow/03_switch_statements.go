package main

import "fmt"

func main() {

	// --- Basic `switch` statement ---
	// Go's `switch` statements do not require an explicit `break` keyword.
	// The first case that matches will be executed, and then the switch is exited.
	fmt.Println("--- Basic Switch ---")
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week.")
	case "Tuesday":
		fmt.Println("Second day of the week.")
	case "Friday":
		fmt.Println("Almost the weekend.")
	default:
		fmt.Println("A regular day.")
	}

	// --- `switch` with multiple cases ---
	fmt.Println("\n--- Switch with multiple cases ---")
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("It's a vowel.")
	case "y":
		fmt.Println("Sometimes a vowel.")
	default:
		fmt.Println("It's a consonant.")
	}

	// --- `switch` with no expression ---
	// This is an alternative to a long `if/else if` chain.
	fmt.Println("\n--- Switch with no expression ---")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: Below C")
	}

	// --- Using `fallthrough` ---
	// `fallthrough` forces execution to continue to the next case, even if the condition is not met.
	// This is rarely used but can be helpful in specific situations.
	fmt.Println("\n--- Switch with `fallthrough` ---")
	number := 2
	switch number {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default case")
	}
}
