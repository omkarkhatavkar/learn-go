package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Go is a statically typed language, which means you must explicitly convert
	// between different types.

	// --- Basic Type Conversion ---
	var myInt int = 100
	var myFloat float64 = float64(myInt) // Explicit conversion from int to float64

	fmt.Println("--- Basic Type Conversion ---")
	fmt.Printf("Original int: %d, Converted float: %f\n", myInt, myFloat)

	// --- String Conversion with `strconv` ---
	// The `strconv` package is used for converting strings to and from other data types.
	
	// int to string
	var intValue int = 123
	stringFromInt := strconv.Itoa(intValue) // Itoa means "Integer to ASCII"
	fmt.Printf("\nOriginal int: %d, Converted string: %q\n", intValue, stringFromInt)

	// string to int
	var stringValue string = "456"
	intFromString, err := strconv.Atoi(stringValue) // Atoi means "ASCII to Integer"
	
	// The `Atoi` function returns two values: the integer and an error.
	// Go's idiomatic error handling uses a two-value return pattern.
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("Original string: %q, Converted int: %d\n", stringValue, intFromString)
	}

	// Let's see what happens with an invalid string
	invalidString := "not-a-number"
	_, err = strconv.Atoi(invalidString) // Use a blank identifier `_` to ignore the return value
	if err != nil {
		fmt.Println("\nHandling invalid conversion:")
		fmt.Println("Error converting 'not-a-number' to int:", err)
	}
}
