package main

import "fmt"

type person struct {
	first  string
	last   string
	favicr []string
}

func main() {
	p1 := person{
		first:  "badamreena",
		last:   "reena",
		favicr: []string{"chocolate", "strawberry", "mango", "coffee", "vanilla"},
	}
	p2 := person{
		first:  "badamsuresh",
		last:   "suresh",
		favicr: []string{"pine apple", "butter scotch"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("-----------------------")

	fmt.Println(p1.favicr)
	fmt.Println(p2.favicr)
	fmt.Println("-----------------------")

	for _, v := range p1.favicr {
		fmt.Println(p1.first, "favorite is", v)
	}
	for _, v := range p2.favicr {
		fmt.Println(p2.first, "favorite is", v)
	}

}
