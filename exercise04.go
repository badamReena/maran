//var for zero value
//short declaration
//multiple intialization
//var when specificity is required
//blank identifier
package main
import "fmt"
var y int
func main(){
	fmt.Println(y)
	z := 4
	fmt.Println(z)
	a , b, c, _, d := 1 , 2, 3, "hello", "world"
	fmt.Println(a , b, c, d)
	var f float32 = 34.7
	fmt.Printf("%v is of type %t\n", f, f)
	g , h, _, i := 45, 2, 67, 9
	fmt.Println(g, h, i)

}
