package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

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

	u3 := user{
		First:   "pooja",
		Last:    "kumar",
		Age:     24,
		Sayings: []string{"i love my friend"},
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)

	//custom sorting
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)

		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("--------------------------------------------")

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)

		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("-------------------------------------------------")

	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)

		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

}
