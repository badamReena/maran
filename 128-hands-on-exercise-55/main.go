package main

import "fmt"

type engine struct {
	electric bool
}

type vehical struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vehical{
		engine: engine{
			electric: true,
		},
		make:  "tata",
		model: "nano",
		doors: 4,
		color: "yellow",
	}
	v2 := vehical{
		engine: engine{
			electric: false,
		},
		make:  "toyota",
		model: "tundra",
		doors: 2,
		color: "black",
	}
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.engine.electric, v1.make, v1.model)
	fmt.Println(v2.engine.electric, v2.make, v2.model)

}
