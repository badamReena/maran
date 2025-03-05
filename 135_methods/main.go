package main

import "fmt"

type person struct {
	first string
}

func (p person) speack() {
	fmt.Println("iam", p.first)
}
func main() {
	p1 := person{
		first: "selva",
	}
	p2 := person{
		first: "suresh",
	}
	p1.speack()
	p2.speack()
}
