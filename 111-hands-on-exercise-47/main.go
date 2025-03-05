package main

import "fmt"

func main() {
	x1 := make([]int, 10)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	x2 := make([]int, 0, 10)
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	fmt.Println("...........................")
	//x1 = append(x1, 12)
	x1[0] = 12
	x2 = append(x2, 13)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))

	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	fmt.Println("............................")

	xs := make([]string, 0, 10)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))

	fmt.Println("...............................")
	xs = append(xs, "a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i], i)
	}
}
