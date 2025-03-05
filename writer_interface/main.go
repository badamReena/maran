package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello,playground")
	fmt.Fprintln(os.Stdout, "hello,playground")

	io.WriteString(os.Stdout, "hello,playgroung")

}
