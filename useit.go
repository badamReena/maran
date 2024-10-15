package furtherexplored

import "fmt"

// this is also "package block" scope
// this is exported because the identifier "fasincating" is capital
func fascinating() {
	fmt.Println("planet speed:", planetspeed)
	planetrotationspeed := 1000
	fmt.Println("planet spinning speed:", planetrotationspeed)

}
 
