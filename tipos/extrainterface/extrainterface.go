package main

import (
	"fmt"
	"math"
)

// type Shape interface {
// 	Area() float64
// }

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	rect := Rectangle{width: 5, height: 3}
	fmt.Println(rect.Area())
	circ := Circle{radius: 3}
	fmt.Println(circ.Area())
}
