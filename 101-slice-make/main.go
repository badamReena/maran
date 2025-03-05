package main

import "fmt"

func main() {
	// xi :=[]int{"a","b","c"}
	// fmt.println(xi) // 0 len 10 capacity
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("xi")
	fmt.Println("----------")
	xi = append(xi, 11, 12, 13, 14, 15)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}
