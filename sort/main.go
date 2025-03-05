package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 3, 8, 78, 34, 26, 96, 55, 12}

	xs := []string{"hello", "welcome", "pooja", "green", "apple"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
