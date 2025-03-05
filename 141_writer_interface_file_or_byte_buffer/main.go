package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {//
	first string
}
func (p person) writeout(w io.Writer) error {//
	_, err := w.Write([]byte(p.first))
	return err
}
func main() {
	p := person{
		first: "reena",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %S", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeout(f)
	p.writeout(&b)
	fmt.Println(b.String())

}
