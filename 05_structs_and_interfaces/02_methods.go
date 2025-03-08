package main

import "fmt"

// Define a `Circle` struct.
type Circle struct {
	Radius float64
}

// A `method` is a function with a `receiver` argument.
// The receiver is specified before the function name.

// This method has a `value receiver`. It operates on a copy of the Circle.
// Changes made to the receiver inside this method will NOT affect the original variable.
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// This method has a `pointer receiver`.
// It operates on the actual Circle instance. Use a pointer receiver when your
// method needs to modify the value of the struct.
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func main() {

	// Create an instance of the `Circle` struct.
	c := Circle{Radius: 10}

	// Calling the `Area` method with a value receiver.
	// Go automatically converts `c` to a pointer `&c` for method calls.
	fmt.Println("--- Methods with Value Receiver ---")
	fmt.Printf("Initial Circle Area: %.2f\n", c.Area())

	// Calling the `Scale` method with a pointer receiver.
	fmt.Println("\n--- Methods with Pointer Receiver ---")
	fmt.Println("Scaling radius by 2...")
	c.Scale(2.0)

	// The original `c` instance has been modified.
	fmt.Printf("Scaled Circle Radius: %.2f\n", c.Radius)
	fmt.Printf("Scaled Circle Area: %.2f\n", c.Area())

	// A pointer to a struct also works for method calls.
	p := &c
	fmt.Println("\n--- Pointer to a struct ---")
	fmt.Printf("Accessing Area from pointer: %.2f\n", p.Area())
}
