package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`//json to go convert//
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"ram","Last":"kumar","Age":25},{"First":"selva","Last":"kumar","Age":19}]`//maeshal output//
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)//888888888888888888//
	if err != nil {//8888888888888888888//
		fmt.Println(err)//888888888888888888888//
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\n person number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
