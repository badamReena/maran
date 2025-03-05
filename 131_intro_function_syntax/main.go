package main

import "fmt"

func main() {
	foo()
	beautifull("soul")
	s := maa("barathi")
	fmt.Println(s)
	s1, n := mama("suresh", 28)
	fmt.Println(s1, n)
}
func foo() {
	fmt.Println("iam from foo")
}
func beautifull(s string) {
	fmt.Println("iam from beautifull soul")
}
func maa(s string) string {
	return fmt.Sprintln("my mom name is", s)
}
func mama(name string, age int) (string, int) {
	age *= 2
	return fmt.Sprintln(name, "is this age was"), age
}
