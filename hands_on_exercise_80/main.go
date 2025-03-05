package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First:   "badam",
		Last:    "reena",
		Age:     25,
		Sayings: []string{"i love u mama"},
	}

	u2 := user{
		First:   "badam",
		Last:    "suresh",
		Age:     28,
		Sayings: []string{"i love u too ammuchu..."},
	}
	users := []user{u1, u2}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("we did something wrong and here the error:", err)
	}
}
