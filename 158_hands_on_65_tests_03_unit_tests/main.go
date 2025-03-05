package main
import "fmt"

func main(){
	fmt.Println(paradise("hawaii"))
}

// paradise returns a string
func paradise(loc string)string{
	return fmt.Sprint("my idea of paradise is", loc)
}