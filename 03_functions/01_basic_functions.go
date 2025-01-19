package main

import "fmt"

// A basic function that takes two integer arguments and returns their sum.
// The return type is specified after the parameter list.
func add(x int, y int) int {
	return x + y
}

// You can also group parameters of the same type.
func subtract(x, y int) int {
	return x - y
}

// A function that doesn't take any arguments and doesn't return any value.
func greet() {
	fmt.Println("Hello from a function!")
}

func main() {
	// Calling the functions and printing the results.
	resultAdd := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", resultAdd)

	resultSubtract := subtract(10, 4)
	fmt.Printf("10 - 4 = %d\n", resultSubtract)

	// Calling a function with no return value.
	greet()
}
