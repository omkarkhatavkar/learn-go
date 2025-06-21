package main

import (
	"context"
	"fmt"
	"time"
)

// `context` is a package used to carry deadlines, cancellation signals, and other request-scoped values
// across API boundaries and between goroutines.

// A goroutine that performs a long-running operation.
func longRunningOperation(ctx context.Context, id int) {
	fmt.Printf("Operation %d: Starting...\n", id)

	select {
	case <-time.After(3 * time.Second):
		// This case will be executed if the operation completes before the context is canceled.
		fmt.Printf("Operation %d: Finished successfully!\n", id)
	case <-ctx.Done():
		// This case is triggered if the context is canceled or its deadline is exceeded.
		fmt.Printf("Operation %d: Canceled! Reason: %s\n", id, ctx.Err())
	}
}

func main() {

	// --- `context` with Timeout ---
	fmt.Println("--- Context with Timeout ---")
	// `context.WithTimeout` creates a new context that is canceled after a certain duration.
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout() // It's good practice to call the cancel function to release resources.

	go longRunningOperation(ctxTimeout, 1)

	// The main goroutine needs to wait to see the result.
	time.Sleep(3 * time.Second)

	// --- `context` with Cancellation ---
	fmt.Println("\n--- Context with Cancellation ---")
	// `context.WithCancel` creates a new context that can be manually canceled.
	ctxCancel, cancelCancel := context.WithCancel(context.Background())

	go longRunningOperation(ctxCancel, 2)
	fmt.Println("Main: Waiting 1 second, then canceling...")
	time.Sleep(1 * time.Second)
	cancelCancel() // Manually cancel the context.

	// The operation will be canceled before it has a chance to finish.
	time.Sleep(2 * time.Second)
}
