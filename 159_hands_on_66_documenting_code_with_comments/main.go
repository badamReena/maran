package main

import "fmt"

func main() {
	fmt.Println(Paradise("hawaii"))
}

// paradise prints out the user's input of paradise as a sentence.
func Paradise(loc string) string {
	return fmt.Sprint("my idea of paradise is", loc)
}
