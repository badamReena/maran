package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)  //value
	fmt.Println(&x) //stored address
	s := "barathi"

	fmt.Printf("%v /t %T /n", &s, &s) //stored address and type
	y := &x
	fmt.Printf("%v /t %T /n", &x, &y) //stored address and type
	fmt.Println(s)                    //value
	fmt.Println(*y)                   //stored value
	fmt.Println(*&x)                  //address stored value
	*y = 34
	fmt.Println(x)  //value
	fmt.Println(*y) //stored value

}
