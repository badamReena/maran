package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("index %v \t value %v\n", i, v)
	}

	fmt.Println("------------")

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("key %v \t value %v\n", k, v)
	}

	fmt.Println("------------")

	age1 := m["James"]
	fmt.Println("the age of bond", age1)

	if v, ok := m["James"]; ok {
		fmt.Println("there is a BOND lookup entry, and Bond's age is", v)
	}
	age2 := m["Q"]
	fmt.Println("that age of Q", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("there is no Q, and here is the zero value of an int", v)
	}

	fmt.Println("-----------")

	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration %v \t total count %v \t x is %v\n", i, c, x)
			c++
		}
	}

}
