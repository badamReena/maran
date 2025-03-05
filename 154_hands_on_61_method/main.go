package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) speak() {
	fmt.Println("my  mom name is", p.first, "and her age is", p.age)
	fmt.Printf("%T/n", p.first, p.age)
}

func main() {
	p1 := Person{
		first: "barathi",
		age:   43,
	}
	p1.speak()
}
