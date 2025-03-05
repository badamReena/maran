package main

import "fmt"

func main() {
	x := 43

	fmt.Println(x)
	fmt.Println(&x)

	s := "selva"

	fmt.Printf("%v /t %T /n", &x, &x)
	fmt.Printf("%v /t %T /n", &s, &s)

}
