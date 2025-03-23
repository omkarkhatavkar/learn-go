package main

import (
	"fmt"
	"time"
)

// A simple function that will be run as a goroutine.
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// A goroutine is a lightweight, independent thread of execution.

	fmt.Println("--- Goroutines Example ---")
	// `go say("world")` starts a new goroutine.
	// The `main` goroutine continues to execute without waiting for the new goroutine to finish.
	go say("world")

	// The `main` goroutine's code runs concurrently with the `say("world")` goroutine.
	say("hello")

	// Notice how the output is interleaved. This is because both goroutines are running at the same time.
	// The program exits when the `main` function finishes, so if `say("world")` takes longer,
	// its output might be cut short.
	
	// To ensure all goroutines complete, you often use `sync.WaitGroup` or channels.
	// We'll see channels in the next example.
}
