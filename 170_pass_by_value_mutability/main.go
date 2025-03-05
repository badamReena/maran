package main

import "fmt"

func intdelta(n *int) {
	*n = 45
}

func slicedelta(ii []int) {
	ii[0] = 99
}

func mapdelta(md map[string]int, s string) {
	md[s] = 32
}

func main() {
	a := 42
	fmt.Println(a)
	intdelta(&a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	slicedelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["selva"] = 28
	fmt.Println(m["selva"])
	mapdelta(m, "selva")
	fmt.Println(m["selva"])
}
