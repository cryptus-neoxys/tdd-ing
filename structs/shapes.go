package main

type Rectangle struct {
	width float64
	height float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.height + rect.width)
}

func Area(rect Rectangle) float64 {
	return rect.width * rect.height
}