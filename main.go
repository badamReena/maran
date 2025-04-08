package main

import (
	"fmt"

	"github.com/reena/go-bookstore/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
