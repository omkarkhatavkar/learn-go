package main

import "fmt"

// A variadic function can accept a variable number of arguments of a specified type.
// The `...` syntax makes the `numbers` parameter a slice of `int`s.
func sumAll(numbers ...int) int {

	// Inside the function, `numbers` is treated as a slice.
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	// We can call the `sumAll` function with different numbers of arguments.
	fmt.Println("Sum of 1, 2, 3:", sumAll(1, 2, 3))
	fmt.Println("Sum of 10, 20, 30, 40, 50:", sumAll(10, 20, 30, 40, 50))

	// You can also pass a slice to a variadic function by using the `...` again.
	// This unpacks the slice into individual arguments.
	fmt.Println("\n--- Passing a Slice ---")
	nums := []int{100, 200, 300}
	fmt.Println("Sum of slice elements:", sumAll(nums...))
}
