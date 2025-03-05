package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)
}

func foo() int {
	return 3
}

func bar() (int, string) {
	return 4, "i love u amma"
}
