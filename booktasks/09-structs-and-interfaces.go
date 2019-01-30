package main

import "math"

func main() {

}

type Shape interface {
	area() float64
	perimeter() float64
}
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, x3, x4 float64
}

func (c Circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r Rectangle) area() float64 {
	return r.x1 * r.x3
}

func (r Rectangle) perimeter() float64 {
	return r.x1 + r.x2 + r.x3 + r.x4
}
