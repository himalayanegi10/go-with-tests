package structsmethodsinterfaces

import (
	"math"
)

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius 
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.base * t.height)
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base float64
	height float64
}

type Shape interface {
	Area() float64
}