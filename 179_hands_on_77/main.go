package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		first: "selva",
	}
	fmt.Println(p)

	p = changename(p, "jayaraman")
	fmt.Println(p)

	changenameptr(&p, "barathi")
	fmt.Println(p)
}

func changename(p person, s string) person {
	p.first = s
	return p
}

func changenameptr(p *person, s string) {
	p.first = s
}
