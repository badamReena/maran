package main

import "fmt"

func main() {
	fmt.Println(printsquare(square, 3))
}

func printsquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number is %v , the square value is %v", x, a)
}

func square(n int) int {
	return n * n
}
