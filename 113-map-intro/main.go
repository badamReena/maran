package main

import "fmt"

func main() {
	am := map[string]int{
		"jam":   46,
		"varun": 16,
		"selva": 34,
	}
	fmt.Println("the age of selva was", am["selva"])
	fmt.Println(am)
	fmt.Printf("#%v\n", am)
	fmt.Println(".......................")
	an := make(map[string]int)
	an["mala"] = 20
	an["suresh"] = 28
	an["reena"] = 25
	fmt.Println(an)
	fmt.Printf("#%v\n", an)
	fmt.Println(len(an))
}
