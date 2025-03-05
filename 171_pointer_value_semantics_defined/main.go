package main

import "fmt"

//valuesemantics

func addone(v int) int {
	return v + 1
}

//pointer semantics
func addonep(v *int) {
	*v += 1
}

func main() {
	//value semantics

	a := 1
	fmt.Println(a)
	fmt.Println(addone(a))

	//pointer semantics
	b := 1
	fmt.Println("value semantics")
	fmt.Println(b)
	addonep(&b)
	fmt.Println(b)
}
