package main

import "fmt"

const (
	_ = iota //c0 == 0
	a
	b
	c
	d
	e
	f
)

func main() {
	//fmt.Println(c0, c1, c2)
	//fmt.Println(c3, c4, c5, c6)
	fmt.Println("%d \t %b\n", 1, 1)
	fmt.Println("%d \t %b\n", 1<<a, 1<<a)
	fmt.Println("%d \t %b\n", 1<<b, 1<<b)
	fmt.Println("%d \t %b\n", 1<<c, 1<<c)
	fmt.Println("%d \t %b\n", 1<<d, 1<<d)
	fmt.Println("%d \t %b\n", 1<<e, 1<<e)
	fmt.Println("%d \t %b\n", 1<<f, 1<<f)
}
