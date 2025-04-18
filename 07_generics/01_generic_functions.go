package main

import "fmt"

// Generics were introduced in Go 1.18. They allow you to write functions and data structures
// that work with any type, while still providing compile-time type safety.

// The `[T any]` syntax declares a type parameter `T`. `any` is a new alias for `interface{}`
// and means `T` can be any type.
func printAny[T any](s T) {
	fmt.Printf("Value: %v, Type: %T\n", s, s)
}

// --- Constraints and Type Sets ---
// You can constrain a type parameter to a specific set of types.
// `comparable` is a pre-defined constraint for types that can be compared with `==` and `!=`.
func compareAndPrint[T comparable](a, b T) {
	fmt.Printf("Are %v and %v equal? %t\n", a, b, a == b)
}

// You can also create your own type sets.
// This `Number` interface is a constraint that allows `int`, `int64`, and `float64`.
type Number interface {
	int | int64 | float64
}

// A generic function that works on any `Number` type.
func add[T Number](a, b T) T {
	return a + b
}

func main() {

	// --- Generic Function Example ---
	fmt.Println("--- Generic Functions ---")
	printAny("Hello, Generics!")
	printAny(123)
	printAny(3.14)

	// --- Using a constrained generic function ---
	fmt.Println("\n--- Constrained Generics ---")
	compareAndPrint("hello", "world")
	compareAndPrint(100, 100)

	// --- Using a custom constraint ---
	fmt.Println("\n--- Custom Constraints ---")
	// The `add` function works for different number types.
	var myInt int = 5
	var myInt64 int64 = 10
	var myFloat64 float64 = 15.5
	
	fmt.Println("Sum of two ints:", add(myInt, 20))
	fmt.Println("Sum of two int64s:", add(myInt64, 25))
	fmt.Println("Sum of two float64s:", add(myFloat64, 30.5))
}
