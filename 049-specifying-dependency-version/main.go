package main

import (
	"fmt"

	"github.com/badamReena/puppy"
)

func main() {
	puppy.From12()
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := puppy.Bark()
	s4 := puppy.Barks()
	fmt.Println(s3)
	fmt.Println(s4)
	//also like this
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}
