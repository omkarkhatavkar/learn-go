package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Go's standard library has a powerful `encoding/json` package for working with JSON.

// We'll use a `struct` to represent our data. The field tags `json:"..."` are crucial.
// They specify the JSON key name for each field.

// A struct for a user.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Email    string `json:"-"` // The `json:"-"` tag ignores this field when encoding or decoding.
}

func main() {

	// --- Encoding (Marshal) a struct to JSON ---
	fmt.Println("--- Encoding Go Struct to JSON ---")
	user := User{
		ID:       1,
		Name:     "Alice",
		IsActive: true,
		Email:    "alice@example.com",
	}

	// `json.Marshal` converts a Go value into a JSON byte slice.
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// The result is a byte slice, so we need to convert it to a string to print it.
	fmt.Printf("JSON data: %s\n", jsonData)

	// --- Decoding (Unmarshal) JSON to a Go struct ---
	fmt.Println("\n--- Decoding JSON to Go Struct ---")
	jsonString := `{"id":2,"name":"Bob","is_active":false}`

	// Create an empty struct to hold the decoded data.
	var user2 User

	// `json.Unmarshal` decodes the JSON data and stores it in the struct.
	// It expects a byte slice, so we convert the string.
	err = json.Unmarshal([]byte(jsonString), &user2)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Printf("Decoded user: %+v\n", user2)
	fmt.Printf("\tID: %d\n", user2.ID)
	fmt.Printf("\tName: %s\n", user2.Name)
	fmt.Printf("\tIs Active: %t\n", user2.IsActive)

	// Note that the `Email` field is the zero value because it was not in the JSON.
	// Also, `user2`'s email is empty because we marked it to be ignored (`json:"-"`) in the struct.
	fmt.Printf("\tEmail (ignored in JSON): %s\n", user2.Email)
}
