package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {
	// Specify the path to your book or text file here
	filePath := "suresh.txt" // Change this to your file path

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read through the file
	scanner := bufio.NewScanner(file)
	wordCount := 0

	// Read through each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into words by whitespace (space, tabs, newlines)
		words := strings.Fields(line)
		wordCount += len(words)
	}

	// Check for any errors during file reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Output the total word count
	fmt.Printf("Total words in the file: %d\n", wordCount)
}
