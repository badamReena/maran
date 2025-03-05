package main

import (
	"log"
	"os"
)

// type person struct {
// first string
//}

// func (p person) writeout(w io.writer) error {
// _, err := w.write([]byte(p.first))
// return err
// }

func main() {
	f, err := os.Create("output.txt")

	if err != nil {
		log.Fatalf("error %s", err)
	}
	//
	defer f.Close()
	//
	s := []byte("hello mama!") //
	_, err = f.Write(s)

	if err != nil {
		log.Fatalf("error %s", err)
	}
}
