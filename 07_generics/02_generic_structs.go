package main

import (
	"fmt"
	"sync"
)

// You can also use generics to create type-safe data structures.

// A generic `Stack` that can hold any type `T`.
// We use a mutex to make it safe for concurrent access.
type Stack[T any] struct {
	items []T
	mu    sync.Mutex
}

// `NewStack` is a constructor for our generic stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// The `Push` method adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
}

// The `Pop` method removes and returns the top item from the stack.
// It returns the item and a boolean indicating if the pop was successful.
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the stack is empty.
	if len(s.items) == 0 {
		var zero T // Return the zero value for type T
		return zero, false
	}

	// Pop the last item from the slice.
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return item, true
}

func main() {

	fmt.Println("--- Generic Stack Example ---")
	// Create a stack of strings.
	stringStack := NewStack[string]()

	stringStack.Push("Apple")
	stringStack.Push("Banana")
	stringStack.Push("Cherry")

	// Pop items and print them.
	fmt.Println("Stack of strings:")
	for {
		item, ok := stringStack.Pop()
		if !ok {
			break // Stack is empty
		}
		fmt.Println(item)
	}

	fmt.Println("\n--- Generic Stack with Integers ---")
	// Create a stack of integers.
	intStack := NewStack[int]()

	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	// Pop items and print them.
	fmt.Println("Stack of integers:")
	for {
		item, ok := intStack.Pop()
		if !ok {
			break
		}
		fmt.Println(item)
	}
}
