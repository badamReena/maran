package main
import "fmt"
func main(){
	/*storing a value of type *string *float *int using print verbs */
	s, i, f := "hello", 78, 34.5
	fmt.Printf("%v is of type %t\n",s ,s ,s)
	fmt.Printf("%v is of type %t\n",i ,i ,i)
	fmt.Printf("%v is of type %t\n",f ,f ,f)

}