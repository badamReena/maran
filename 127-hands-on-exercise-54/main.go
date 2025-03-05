package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	favicr []string
}

func main() {
	p1 := person{
		first:  "badam",
		last:   "reena",
		favicr: []string{"chocolate", "strawberry", "mango", "coffee", "vanilla"},
	}
	p2 := person{
		first:  "badam",
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
	fmt.Println("-----------------------------")
	m := map[string]person{ //
		p1.last: p1, //
		p2.last: p2, //
	}
	for _, v := range m { //
		fmt.Println(v)                //
		for _, v2 := range v.favicr { //
			fmt.Println(v.first, v.last, v2) //
		}
	}

}
