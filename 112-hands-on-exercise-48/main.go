package main

import "fmt"

func main() {
	rj := []string{"reena", "jayaraman", "i love my dad"}
	sr := []string{"suresh", "reena", "i love my hubby"}
	xp := [][]string{rj, sr}

	for i, v := range xp {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
