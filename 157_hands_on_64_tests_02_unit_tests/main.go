package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", domath)

	x := domath(42, 16, add)
	fmt.Println(x)
	
	y := domath(42, 16, subtract)
	fmt.Println(y)
}

func domath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
