/*write a program that uses print verbs to show the following numbers
*747=decimal
*911=binary
*90210=hexadecimal*/
package main

import "fmt"

func main() {
	d, b, h := 747, 911, 90210
	fmt.Printf("%d \t\t %b \t\t %#X\n", d, d, d)
	fmt.Printf("%d \t\t %b \t\t %#X\n", b, b, b)
	fmt.Printf("%d \tt %b \t\t %#X\n", h, h, h)

}
