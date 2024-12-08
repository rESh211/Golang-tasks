package abstract

import (
	"fmt"
	"math"
)

type Circle struct {
	Shape
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	Shape
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Shape interface {
	Area() float64
}

func PrintArea(shape Shape) {
	fmt.Println(shape.Area())
}

func Run() {
	circle := &Circle{radius: 1}
	rectangle := &Rectangle{width: 5, height: 3}
	PrintArea(circle)
	PrintArea(rectangle)
}
