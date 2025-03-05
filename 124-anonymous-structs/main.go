package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "badam",
		last:  "reena",
		age:   25,
	}

	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "badam",
		last:  "suresh",
		age:   28,
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
