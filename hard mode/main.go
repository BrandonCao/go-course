package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if the file path is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		return
	}

	// Get the file path from the command line argument
	filePath := os.Args[1]

	// Read the contents of the file
	content, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the contents to the terminal
	io.Copy(os.Stdout, content)
}
