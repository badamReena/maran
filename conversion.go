package main
import ("fmt")
func main(){
	x := 42
	y := 42.0
	fmt.Printf("%v of is type %t\n", x, x)
	fmt.Printf("%v of is type %t\n", y, y)
	var m float32 = 43.742
	fmt.Printf("%v of is type %t\n", m, m)
	y = float64(m)
	fmt.Printf("%v of is type %t\n", y, y)

}