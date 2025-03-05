package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := bytes.NewBufferString("welcome reena") //
	fmt.Println(b.String())

	b.WriteString("love u suresh mama")
	fmt.Println(b.String())

	//b.Reset()

	b.WriteString("mama")
	fmt.Println(b.String())

	b.Write([]byte("come on...come on...my love..!")) //
	fmt.Println(b.String())
}

// o/p
/*welcome reena
welcome reenalove u suresh mama
welcome reenalove u suresh mamamama
welcome reenalove u suresh mamamamacome on...come on...my love..!*/
