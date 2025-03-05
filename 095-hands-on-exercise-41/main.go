package main

import "fmt"

func main() {
	sr := [...]string{"banana", "ghee", "nuts", "graps", "apple", "milk", "milk shake", "ice cream", "stabery", "watermellon"}
	fmt.Println(sr)
	fmt.Printf("%T\n", sr)

	for i, v := range sr {
		fmt.Printf("%v  - %v\n", i, v)
	}
}
