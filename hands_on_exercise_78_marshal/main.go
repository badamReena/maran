package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Last  string
	Age   int
}

func main() {
	u1 := user{
		First: "jeeva",
		Last:  "karan",
		Age:   34,
	}

	u2 := user{
		First: "barathi",
		Last:  "jayaraman",
		Age:   54,
	}

	u3 := user{
		First: "kala",
		Last:  "vathi",
		Age:   24,
	}

	user := []user{u1, u2, u3}
	fmt.Println(user)

	bs, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
