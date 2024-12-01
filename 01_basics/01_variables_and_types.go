package main

import "fmt"

func main() {

	// Variable declaration with var. This is the more formal way.
	// Go will automatically assign the zero value if no initial value is provided.
	var integerVar int
	var stringVar string = "Hello, Go!"
	var boolVar bool
	var floatVar float64 = 3.14

	fmt.Println("--- Variable Declaration ---")
	fmt.Println("Integer variable (zero value):", integerVar)
	fmt.Println("String variable:", stringVar)
	fmt.Println("Boolean variable (zero value):", boolVar)
	fmt.Println("Float variable:", floatVar)

	// Short variable declaration with :=. This is the most common way to declare variables.
	// Go infers the type from the initial value.
	inferredInt := 42
	inferredString := "Go is fun"
	
	fmt.Println("\n--- Short Declaration ---")
	fmt.Println("Inferred integer:", inferredInt)
	fmt.Println("Inferred string:", inferredString)

	// Constants are values that cannot be changed after declaration.
	const Pi float64 = 3.14159
	const Epsilon = 0.0001 // Type can be inferred

	fmt.Println("\n--- Constants ---")
	fmt.Println("Pi constant:", Pi)
	fmt.Println("Epsilon constant:", Epsilon)
}
