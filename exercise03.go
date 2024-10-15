package main

import "fmt"

type bytesize int

const ( //iota = 0
	kb bytesize = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
)

func main() {

	fmt.Printf("%d \t \t \t %b\n", kb, kb)
	fmt.Printf("%d \t \t %b\n", mb, mb)
	fmt.Printf("%d \t \t  %b\n", gb, gb)
	fmt.Printf("%d \t \t  %b\n", tb, tb)
	fmt.Printf("%d \t  %b\n", pb, pb)
	fmt.Printf("%d \t  %b\n", eb, eb)
}
