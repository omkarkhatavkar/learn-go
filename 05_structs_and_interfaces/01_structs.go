package main

import "fmt"

// A `struct` is a typed collection of fields. It's similar to a class or an object
// in other languages, but it only contains data.

// Define a `Person` struct with three fields.
type Person struct {
	Name string
	Age  int
	City string
}

// You can also embed a struct within another struct.
// This promotes the fields of the embedded struct to the outer struct.
type Employee struct {
	// `Person` is an anonymous field, which means its fields are "promoted"
	// and can be accessed directly from `Employee`.
	Person
	ID int
	Position string
}

func main() {

	// --- Creating and Initializing a struct ---
	// The `var` keyword creates a struct with zero values.
	var p1 Person
	fmt.Println("--- Zero-value struct ---")
	fmt.Println(p1)

	// Field names can be used to initialize the struct.
	// This is the preferred way as it's more readable.
	p2 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Println("\n--- Initialized struct ---")
	fmt.Println(p2)

	// You can omit field names, but it's less readable and dangerous if fields are re-ordered.
	// p3 := Person{"Bob", 25, "London"}

	// --- Accessing fields ---
	fmt.Println("\n--- Accessing Fields ---")
	fmt.Printf("Name: %s, Age: %d\n", p2.Name, p2.Age)

	// You can also change a field's value.
	p2.Age = 31
	fmt.Println("Updated age:", p2.Age)

	// --- Embedded structs ---
	fmt.Println("\n--- Embedded Structs ---")
	e := Employee{
		Person: Person{
			Name: "Carol",
			Age:  45,
			City: "Paris",
		},
		ID:       101,
		Position: "Engineer",
	}

	// Accessing promoted fields directly from the outer struct.
	fmt.Printf("Employee Name: %s, Position: %s\n", e.Name, e.Position)

	// You can also access the embedded struct itself.
	fmt.Printf("Employee's Person struct: %+v\n", e.Person)
}
