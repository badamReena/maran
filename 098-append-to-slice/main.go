//func append (slice[] type, elems...type)[]type

package main

import "fmt"

func main() {
	xi := []int{46, 47, 48}
	fmt.Println(xi)
	fmt.Println("-------")
	// variadic parameter
	xi = append(xi, 45, 47, 48, 49, 50)
	fmt.Println(xi)
	fmt.Println("-------")
}
