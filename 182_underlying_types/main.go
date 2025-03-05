package main

import "fmt"

func addi(a, b int) int {
	return a + b
}

func addf(a, b float64) float64 {
	return a + b
}

type mynumber interface {
	~int | ~float64
}

func addt[t mynumber](a, b t) t {
	return a + b
}

type myalias int

func main() {
	var n myalias = 43

	fmt.Println(addi(1, 2))
	fmt.Println(addf(1.2, 2.2))

	fmt.Println(addt(n, 2))
	fmt.Println(addt(1.2, 2.2))
}
