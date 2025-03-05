package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretagent struct {
	person
	ltk   bool
	first string
}

func main() {
	sa1 := secretagent{
		person: person{
			first: "badam",
			last:  "reena",
			age:   25,
		},
		first: "kala",
		ltk:   true,
	}
	p2 := person{
		first: "badam",
		last:  "suresh",
		age:   28,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Printf("%T \t %#v\n", sa1, sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person)
}
