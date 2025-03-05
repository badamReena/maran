package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"jeeva","Last":"karan","Age":34},{"First":"barathi","Last":"jayaraman","Age":54},{"First":"kala","Last":"vathi","Age":24}]`
	fmt.Println(s)

	var people []person
	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, person := range people {

		fmt.Println("Person #", i)

		fmt.Println("\t", person.First, person.Last, person.Age)

		//for _, Age := range person.Age{
		//	fmt.Println("\t\t",Age)
		//}
	}
}
