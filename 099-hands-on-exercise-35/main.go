package main
import "fmt"
func main(){
	x := "james bond"
	if x == "james bond" {
		fmt.Println(x)
	}
	y := "moneypenny"
	if y == "moneypenny" {
		fmt.Println(y)
	} else if y == "james bond" {
		fmt.Println("bonddonbondonbond", y)
	} else {
		fmt.Println("neither")
	}

	switch {
	case false:
		fmt.Println("doesn't print")
	case true:
		fmt.Println("prints")
	case true:
		fmt.Println("doesn't print")
	}

	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingssuit flying":
		fmt.Println("living dangerously")
	}
}