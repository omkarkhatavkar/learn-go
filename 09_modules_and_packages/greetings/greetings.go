package greetings

import "fmt"

// Go's conventions for public and private access are based on capitalization.
// An identifier (function, variable, struct, etc.) that starts with a capital letter is **exported**
// and can be accessed from other packages.
// An identifier that starts with a lowercase letter is **unexported** (private) and can only
// be used within its own package.

// Hello returns a greeting for a given name.
func Hello(name string) string {
	// `fmt.Sprintf` formats a string without printing it.
	message := fmt.Sprintf("Hi, %s. Welcome!", name)
	return message
}

// `goodbye` is a private function in this package.
func goodbye() string {
	return "Goodbye!"
}
