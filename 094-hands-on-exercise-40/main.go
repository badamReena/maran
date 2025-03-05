package main

import "fmt"

func main() {
	a := [...]string{"milk", "cake", "tomato", "potato", "jam", "bread", "dish wash", "water can", "oil", "steel", "coper", "bronze"}
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	fmt.Println(len(a))

}
