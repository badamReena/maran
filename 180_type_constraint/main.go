package main

import "fmt"

func addi(a, b int) int {
	return a + b
}

func addf(a, b float64) float64 {
	return a + b
}

func addt[t int | float64](a, b t) t {
	return a + b
}

func main() {
	fmt.Println(addi(1, 2))
	fmt.Println(addf(1.2, 2.2))

	fmt.Println(addt(1, 2))
	fmt.Println(addt(1.2, 2.2))
}
