package main

import "fmt"

func main() {
	sr := [...]string{"banana", "ghee", "nuts", "graps", "apple", "milk", "milk shake", "ice cream", "stabery", "watermellon"}
	fmt.Println(sr)
	fmt.Printf("%T\n", sr)
	//blank identifier to not use a returned value
	for _, v := range sr {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("---------")
	fmt.Println(sr[0])
	fmt.Println(sr[1])
	fmt.Println(sr[2])
	//doesn't work
	//fmt.println(sr[3])
	fmt.Println(len(sr))
	for i := 0; i < len(sr); i++ {
		fmt.Println(sr[i])
	}
}
