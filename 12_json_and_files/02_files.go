package main

import (
	"fmt"
	"os"
	"log"
)

// The `os` package provides a platform-independent interface to operating system functionality,
// including file I/O.

func main() {

	// --- Writing a file ---
	fmt.Println("--- Writing a File ---")
	// `os.WriteFile` is a simple way to write data to a file.
	// It takes the filename, the data as a byte slice, and the file permissions.
	data := []byte("Hello, Go file!\nThis is a test file.\n")
	err := os.WriteFile("test.txt", data, 0644) // 0644 is a common permission setting
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
	fmt.Println("Successfully wrote to test.txt.")

	// --- Reading a file ---
	fmt.Println("\n--- Reading a File ---")
	// `os.ReadFile` reads the entire file into a byte slice.
	fileData, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert the byte slice to a string to print.
	fmt.Println("File content:")
	fmt.Println(string(fileData))

	// --- Deleting a file ---
	fmt.Println("\n--- Deleting a File ---")
	err = os.Remove("test.txt")
	if err != nil {
		log.Fatalf("Error deleting file: %v", err)
	}
	fmt.Println("Successfully deleted test.txt.")

	// --- More advanced file I/O ---
	// For larger files or streaming, you would use `os.OpenFile` and `bufio`.
	// This provides more control and is more efficient.
	fmt.Println("\n--- Advanced File I/O ---")
	file, err := os.Create("advanced.txt")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close() // `defer` ensures the file is closed at the end of the function.

	_, err = file.WriteString("Writing a line to an advanced file.\n")
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
	fmt.Println("Successfully wrote to advanced.txt.")
	
	// Clean up the file
	defer os.Remove("advanced.txt")
}
