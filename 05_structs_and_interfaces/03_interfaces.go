package main

import "fmt"

// An `interface` in Go is a set of method signatures.
// A type implicitly implements an interface if it has all the methods defined by the interface.

// Define a `Shape` interface.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Now, define a `Rectangle` struct.
type Rectangle struct {
	Width, Height float64
}

// The `Rectangle` type implements the `Shape` interface implicitly by providing
// `Area()` and `Perimeter()` methods with the correct signatures.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Define a `Circle` struct. It will also implement the `Shape` interface.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// A function that accepts any type that satisfies the `Shape` interface.
func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {

	// Create instances of our structs.
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}

	// We can pass both `Rectangle` and `Circle` to the same function because
	// they both implement the `Shape` interface.
	fmt.Println("--- Rectangle Details ---")
	printShapeDetails(r)

	fmt.Println("\n--- Circle Details ---")
	printShapeDetails(c)

	// --- The Empty Interface `interface{}` or `any` ---
	// The empty interface is satisfied by any type, which makes it useful for handling
	// values of unknown type. It is now aliased to `any`.
	fmt.Println("\n--- The Empty Interface ---")
	var i interface{} // or var i any
	i = "Hello"
	fmt.Printf("Value: %v, Type: %T\n", i, i)

	i = 42
	fmt.Printf("Value: %v, Type: %T\n", i, i)

	// --- Type Assertion ---
	// To access the underlying value of an interface, you must use a type assertion.
	// `value, ok := i.(<type>)` is the idiomatic way.
	fmt.Println("\n--- Type Assertion ---")	
	val, ok := i.(int)
	if ok {
		fmt.Printf("Assertion successful! Value: %d\n", val)
	}

	valFloat, ok := i.(float64)
	if !ok {
		fmt.Println("Assertion to float64 failed.")
	}
}
