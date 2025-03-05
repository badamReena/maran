package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []string{"reena", "amaran", "kaviya", "pooja", "bhavani"}
	xs := []int{23, 56, 34, 23, 80, 12, 3, 6, 8, 4}

	fmt.Println(xi)
	sort.Strings(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Ints(xs)
	fmt.Println(xs)
}
