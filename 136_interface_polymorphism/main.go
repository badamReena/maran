package main

import "fmt"

type person struct { //
	first string
}

type secretagent struct { //
	person
	ltk bool
}

func (p person) speack() {
	fmt.Println("iam", p.first)
}
func (sa secretagent) speack() {
	fmt.Println("iam a secretagent", sa.first)
}

type human interface { //
	speack()
}

func saysomething(h human) {
	h.speack()
}

func main() {
	sa1 := secretagent{
		person: person{
			first: "reena",
		},
		ltk: true,
	}
	p2 := person{
		first: "suresh",
	}
	// sa1.speack()
	// p2.speack()
	saysomething(sa1)
	saysomething(p2)

}
