package main

import "fmt"

func main() {

	// Maps are a built-in key-value data structure, also known as a hash map or dictionary.

	// --- Declaring and initializing a map ---
	// The syntax is `map[keyType]valueType`.
	// Use `make` to create an empty map before using it.
	studentGrades := make(map[string]int)

	// --- Adding and updating entries ---
	studentGrades["Alice"] = 95
	studentGrades["Bob"] = 88
	studentGrades["Charlie"] = 92

	fmt.Println("--- Initial Map ---")
	fmt.Println("Student grades:", studentGrades)

	// --- Accessing a value ---
	// Accessing a key that doesn't exist returns the zero value for the value type (e.g., 0 for int).
	fmt.Println("\n--- Accessing Values ---")
	fmt.Println("Alice's grade:", studentGrades["Alice"])
	fmt.Println("Dave's grade (doesn't exist):", studentGrades["Dave"])

	// --- Checking for key existence with the 'comma ok' idiom ---
	// This is the idiomatic way to check if a key exists in a map.
	// The second return value `ok` will be `true` if the key exists, `false` otherwise.
	grade, ok := studentGrades["Bob"]
	fmt.Println("\n--- 'Comma Ok' Idiom ---")
	fmt.Printf("Bob's grade: %d, exists: %t\n", grade, ok)

	grade, ok = studentGrades["Eve"]
	fmt.Printf("Eve's grade: %d, exists: %t\n", grade, ok)

	// --- Deleting an entry ---
	delete(studentGrades, "Charlie")
	fmt.Println("\n--- Deleting an Entry ---")
	fmt.Println("Map after deleting Charlie:", studentGrades)

	// --- Iterating over a map ---
	// Use a `for-range` loop to iterate over key-value pairs.
	fmt.Println("\n--- Iterating over Map ---")
	for name, grade := range studentGrades {
		fmt.Printf("Name: %s, Grade: %d\n", name, grade)
	}

	// You can also declare and initialize a map in a single line.
	countries := map[string]string{
		"USA": "United States",
		"FR":  "France",
		"JP":  "Japan",
	}
	fmt.Println("\n--- Initialized Map ---")
	fmt.Println("Countries:", countries)
}
