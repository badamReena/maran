package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width //importan//
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2) //importan//
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}
func main() {
	c1 := circle{
		radius: 6,
	}
	s1 := square{
		length: 4,
		width:  3,
	}
	fmt.Println(info(c1))
	fmt.Println(info(s1))
}
