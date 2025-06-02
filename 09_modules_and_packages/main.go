package main

import (
	"fmt"
	"learning/modules/greetings" // This imports the `greetings` package
)

// Go programs are organized into packages.
// The `main` package is a special package that defines an executable program.

func main() {
	// We can now call a function from the `greetings` package.
	message := greetings.Hello("Go Developer")
	fmt.Println(message)
}
