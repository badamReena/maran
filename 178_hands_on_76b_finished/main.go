package main

import (
	"fmt"
)

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("my name is ", d.first, "and iam walking rightnow.")
}

func (d dog) run() {
	fmt.Println("my name is ", d.first, "and iam running rightnow.")
}

type youngin interface {
	walk()
	run()
}

func yougrun(y youngin) {
	y.run()
	y.walk()
}

func main() {
	d1 := dog{"guna"}
	yougrun(d1)
	//youngrun(d1) //doesn't run

	d2 := dog{"deepa"}
	yougrun(d2)
	
	d2 = d2.changename("rover")
	yougrun(d2)
}

func (d dog) changename(s string) dog {
	d.first = s
	return d
}
