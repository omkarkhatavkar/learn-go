package main

import "fmt"

func main() {

	// Arrays in Go are fixed-size. The size is part of the type.
	// Here we declare an array of 3 strings.
	var colors [3]string

	// Assigning values to array elements by index.
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("--- Basic Array --- ")
	fmt.Println("My colors:", colors)
	fmt.Println("Number of colors:", len(colors))

	// You can also declare and initialize an array in a single line.
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("\n--- Initialized Array --- ")
	fmt.Println("My numbers:", numbers)

	// Using the ellipsis (...) to have the compiler count the elements for you.
	moreNumbers := [...]int{10, 20, 30, 40, 50, 60}
	fmt.Println("\n--- Ellipsis Array --- ")
	fmt.Println("More numbers:", moreNumbers)
	fmt.Println("Number of more numbers:", len(moreNumbers))

	// Iterating over an array using a `for` loop with `range`.
	fmt.Println("\n--- Iterating with range --- ")
	// `i` is the index and `color` is the value.
	for i, color := range colors {
		fmt.Printf("Index: %d, Value: %s\n", i, color)
	}

	// To iterate over values only, use the blank identifier `_` for the index.
	fmt.Println("\n--- Iterating over values only --- ")
	for _, number := range numbers {
		fmt.Printf("Value: %d\n", number)
	}
}
