package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int//
		favdrinks []string//
	}{
		first: "reena",
		friends: map[string]int{
			"poo":   24,
			"meenu": 24,
			"nivi":  24,
		},
		favdrinks: []string{
			"rosemilk",
			"badammilk",
		},
	}
	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends - ", k, v)

		for _, v  := range p1.favdrinks {
			fmt.Println(p1.first, "-drinks - ", v)
		}
	}
}
