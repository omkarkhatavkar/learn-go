package main

import (
	"fmt"
	"time"
)

// The `select` statement lets a goroutine wait on multiple communication operations.
// A `select` blocks until one of its cases can run, then it executes that case.

func main() {
	// Create two channels.
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Start two goroutines to send a message to each channel after a delay.
	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "two"
	}()

	// We'll use a `select` statement to listen for messages from both channels.
	fmt.Println("--- `select` Statement Example ---")
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received from channel 1:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received from channel 2:", msg2)
		}
	}

	// The first message will be from `channel1` because its goroutine finishes first.
	// The `select` statement will then block until `channel2` sends its message.

	// --- `select` with a `default` case ---
	// A `default` case in a `select` statement allows for non-blocking communication.
	// If no other case is ready, the `default` case executes immediately.
	fmt.Println("\n--- `select` with `default` ---")
	msg := ""
	select {
	case msg = <-channel1:
		fmt.Println("Received:", msg)
	case msg = <-channel2:
		fmt.Println("Received:", msg)
	default:
		// This will execute immediately because both channels are empty at this point.
		fmt.Println("No messages received, running default.")
	}
}
