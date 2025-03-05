package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "badam",
		last:  "reena",
		age:   25,
	}
	p2 := person{
		first: "badam",
		last:  "suresh",
		age:   28,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T\t %#v\n", p2, p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
