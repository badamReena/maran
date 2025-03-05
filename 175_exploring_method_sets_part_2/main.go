package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("my name is ", d.first, "and iam walking.")
}

func (d *dog) run() {
	d.first = "selva"
	fmt.Println("my name is ", d, d.first, "and iam running in the road.")
}

type youngin interface {
	walk()
	run()
}

func youngrun(y youngin) {
	y.run()
}

func main() {

	d1 := dog{"mara"}
	d1.walk()
	d1.run()
	// youngrun(d1)

	d2 := &dog{"sriram"}
	d2.walk()
	d2.run()
	youngrun(d2)
}
