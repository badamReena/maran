package main

import "fmt"

var x = 40
const y = 42

func main() {
	fmt.Printf("the value of x is %v and the type of x is %T\n", x, x)
	fmt.Printf("the value of x is %v and the type of x is %T\n", y, y)
	z := 41
	fmt.Printf("the value of x is %v and the type of x is %T\n", z, z)

}
