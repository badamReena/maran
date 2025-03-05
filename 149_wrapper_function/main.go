package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readfile("poem.txt")
	if err != nil {
		log.Fatalf("error in main in readfile: %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}
func readfile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return xb, nil
}
