package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {
	// Ask the user for the file path
	fmt.Print("Enter the file path: ")
	var filePath string
	fmt.Scanln(&filePath)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	wordCount := 0

	// Read through the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into words by spaces and other non-alphanumeric characters
		words := strings.Fields(line) // splits by whitespace (spaces, tabs, newlines)
		wordCount += len(words)
	}

	// Check for any error that might have occurred while reading the file
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the total word count
	fmt.Printf("Total words in the file: %d\n", wordCount)
}
