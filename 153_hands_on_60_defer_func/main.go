package main

import "fmt"

func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)
}

/* deferred functions run in LIFO order
last in first out */
