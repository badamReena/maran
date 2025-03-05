package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addi(a, b int) int {
	return a + b
}

func addf(a, b float64) float64 {
	return a + b
}

type mynumbers interface {
	//constraints.integer | constraints.float
	//constraints.Integer
	constraints.Integer | constraints.Float
}

func addt[t mynumbers](a, b t) t {
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
