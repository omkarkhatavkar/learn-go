package main

import (
	"fmt"
	"time"
)

// A channel is a typed conduit through which you can send and receive values
// with a channel operator, `<-`.

func sender(ch chan string) {
	fmt.Println("Sender: Sending messages...")
	time.Sleep(1 * time.Second)
	ch <- "message 1" // Send a string to the channel
	ch <- "message 2" // Send another string
	ch <- "message 3"
	fmt.Println("Sender: Finished sending.")
	close(ch) // Close the channel to signal no more values will be sent
}

func receiver(ch chan string) {
	fmt.Println("Receiver: Waiting for messages...")
	// The `range` loop iterates over values received from the channel until it's closed.
	for msg := range ch {
		fmt.Printf("Receiver: Received %s\n", msg)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Receiver: Channel is closed, finished receiving.")
}

func main() {

	// --- Unbuffered Channel ---
	// An unbuffered channel has a capacity of zero. It blocks until both a sender and a receiver are ready.
	fmt.Println("--- Unbuffered Channel Example ---")
	unbufferedChan := make(chan string)

	// Start the sender and receiver goroutines.
	go sender(unbufferedChan)
	go receiver(unbufferedChan)

	// To prevent the main goroutine from exiting prematurely, we'll wait a bit.
	// In real applications, you would use a `sync.WaitGroup` or another channel for synchronization.
	time.Sleep(4 * time.Second)

	// --- Buffered Channel ---
	// A buffered channel has a capacity. It only blocks when the buffer is full (on send)
	// or empty (on receive).
	fmt.Println("\n--- Buffered Channel Example ---")
	bufferedChan := make(chan string, 2) // Create a buffered channel with a capacity of 2

	bufferedChan <- "buffered msg 1" // Send a message (won't block yet)
	bufferedChan <- "buffered msg 2" // Send a second message (won't block yet)
	fmt.Println("Sent two messages to buffered channel. Buffer is full.")

	// The next send will block until a receiver takes a value.
	go func() {
		fmt.Println("Attempting to send a third message (will block)...")
		bufferedChan <- "buffered msg 3"
		fmt.Println("Successfully sent third message.")
	}()

	time.Sleep(1 * time.Second) // Let the goroutine attempt to send

	fmt.Println("Receiving messages from buffered channel...")
	fmt.Printf("Received: %s\n", <-bufferedChan)
	fmt.Printf("Received: %s\n", <-bufferedChan)

	time.Sleep(1 * time.Second) // Now the blocked goroutine can send its message
	fmt.Printf("Received: %s\n", <-bufferedChan)
	
	fmt.Println("Closing channel.")
	close(bufferedChan)
}
