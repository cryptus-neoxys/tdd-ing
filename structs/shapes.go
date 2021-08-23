package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.height + rect.width)
}

func Area(rect Rectangle) float64 {
	return rect.width * rect.height
}
