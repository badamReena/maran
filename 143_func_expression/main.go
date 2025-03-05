package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("anonymous func ran")
	}
	y := func(s string) {
		fmt.Println("this is an anonymous func showing my name", s)
	}
	x()
	y("suresh")
}
func foo() {
	fmt.Println("welcome foo ran")
}
