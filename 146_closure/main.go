package main

import "fmt"

func main() {

	s := incrementor()
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())

	r := incrementor()
	fmt.Println(r())
	fmt.Println(r())
	fmt.Println(r())
	fmt.Println(r())
	fmt.Println(r())
}
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
