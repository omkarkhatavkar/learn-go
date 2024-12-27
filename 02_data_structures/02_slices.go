package main

import "fmt"

func main() {

	// Slices are a powerful and common data structure in Go.
	// Unlike arrays, slices are dynamic and can grow or shrink.

	// --- Declaring a slice ---
	// The `[]` syntax without a number creates a slice.
	var fruits []string

	fmt.Println("--- Initial Slice ---")
	fmt.Println("Initial fruits slice:", fruits)
	fmt.Println("Length:", len(fruits))
	fmt.Println("Capacity:", cap(fruits))

	// --- Appending to a slice ---
	// Use the built-in `append` function to add elements.
	// `append` returns a new slice, so you must re-assign it.

	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Cherry")

	fmt.Println("\n--- Appended Slice ---")
	fmt.Println("Fruits after appending:", fruits)
	fmt.Println("Length:", len(fruits))
	fmt.Println("Capacity:", cap(fruits))

	// --- Declaring and initializing a slice ---
	animals := []string{"Dog", "Cat", "Bird"}
	fmt.Println("\n--- Initialized Slice ---")
	fmt.Println("Initial animals slice:", animals)

	// --- Slicing an existing slice ---
	// Slicing creates a new slice that references a portion of the original array's memory.
	// `[low:high]` -> `low` is inclusive, `high` is exclusive.
	fmt.Println("\n--- Slicing ---")
	someFruits := fruits[1:3] // Get elements from index 1 (inclusive) to 3 (exclusive)
	fmt.Println("Slice of fruits from index 1 to 2:", someFruits)
	
	// Modifying the new slice will affect the original, because they share the same underlying array.
	someFruits[0] = "Grapes" // This modifies "Banana" to "Grapes" in the `fruits` slice
	fmt.Println("\n--- Slicing Mutability ---")
	fmt.Println("Original fruits after mutation:", fruits)

	// --- Using `make` to create a slice ---
	// `make` allows you to pre-allocate memory, which can be more efficient.
	// `make([]type, length, capacity)`
	numbers := make([]int, 5, 10)
	fmt.Println("\n--- Slice with `make` ---")
	fmt.Println("Numbers slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}
