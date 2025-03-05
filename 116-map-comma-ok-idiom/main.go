package main

import "fmt"

func main() {
	am := map[string]int{
		"jam":   46,
		"varun": 16,
		"selva": 34,
	}
	fmt.Println("the age of selva was", am["selva"])
	fmt.Println(am)
	fmt.Printf("#%v\n", am)
	fmt.Println(".......................")
	an := make(map[string]int)
	an["mala"] = 20
	an["suresh"] = 28
	an["reena"] = 25
	fmt.Println(an)
	fmt.Printf("#%v\n", an)
	fmt.Println(len(an))
	delete(an, "reena")
	fmt.Println(".....accessing keys that don't exist")
	delete(an, "reena")      // won't panic
	fmt.Println(an["reena"]) // won't panic
	fmt.Println("....................")
	//v, ok := an["reena"]
	//if ok {
	//	fmt.Println("the value prints", v)
	//} else {
	//fmt.Println("key didn't exist")
	//}
	if v, ok := an["mala"]; !ok {
		fmt.Println("key didn't exist")
	} else {
		fmt.Println("the value prints", v)
	}
}
