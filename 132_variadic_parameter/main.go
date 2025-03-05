package main

import "fmt"

func main() {
	x := sum("reena", 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("the sum of", x)

}

//func(r reciver)identifier(p parameters)(returns(s)){ code}
func sum(s string, ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
