package main

import (
	"fmt"

	"time"
)

func main() {
	timefunc(dowork)
}

func dowork() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func timefunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
