package main
import "fmt"
func main(){
	// x can be accessed here
	fmt.Fprintln(x)

	// x can be acessed here
	printMe()
	//y does not exist here yet
	// so this won't work
	// fmt.println(y)

	//y is in "block scope"
	y := 43
	fmt.println(y)

	p1 := person{
		"james",
	}
	p1.sayhello()

	// variable "shadowing"
	x := 32
	fmt.Println(x)

	//package block x is still the same
	printMe()
	
	furtherexplored.fascinating()

	 //also good to know

	 if z := 82; z >50 {
		fmt.Println(z)
	 }
//you can't access z here
// fmt.println(z)
/*
take a look:at the "comma ok idiom"
https://go.dev/doc/effective_go#map3 */

}
func printMe(){
	// x can be accesed here
	fmt.Println(x)

}
// type person is in "package block" scope
type person struct{
	first string
}
// the method sayhello
// which is attached to values of type person
// is in "package block" scope
func (p person) sayhello() {
	fmt.Println("hi, my name is",p.first)
}