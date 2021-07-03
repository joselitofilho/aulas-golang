package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Shape interface {
	Area() float64
}

func main() {
	circle := Circle{Radius: 10}
	rectangle := Rectangle{Height: 10, Width: 5}
	fmt.Printf("circle = %f / rectangle = %f\n", circle.Area(), rectangle.Area())

	var shape Shape
	shape = circle
	fmt.Println("   shape circle =", shape.Area())
	shape = rectangle
	fmt.Println("shape rectangle =", shape.Area())
}
