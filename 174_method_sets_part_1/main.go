package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("my name is ", d.first, " and iam walking in the road.")
}

func (d *dog) run() {
	d.first = "selva"
	fmt.Println("my name is", d.first, "and iam running in the road")
}

func main() {
	d1 := dog{"ram"}
	d1.walk()
	d1.run()

	d2 := &dog{"raj"}
	d2.walk()
	d2.run()
}
