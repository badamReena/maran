package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct { //
	title string
}

func (b book) string() string {
	return fmt.Sprint("the tittle of the book", b.title)
}

type count int //

func (c count) string() string {
	return fmt.Sprint("this is the number", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "my love story",
	}

	var n count = 42

	log.Println(b)
	log.Println(n)
}
