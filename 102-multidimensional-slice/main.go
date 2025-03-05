package main

import "fmt"

func main() {
	JB := []string{"JAMES", "BOND", "MARTINI", "CHOCOLATE"}
	JM := []string{"JENNY", "MONEPENNY", "GUINESS", "WOLVERINE TRACKS"}
	fmt.Println(JB)
	fmt.Println(JM)
	xxs := [][]string{JB, JM}
	fmt.Println(xxs)
}
