package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	base   float64
	height float64
}

func main() {
	square := square{
		base:   10,
		height: 5,
	}

	triangle := triangle{
		base:   5,
		height: 20,
	}

	printArea(square)
	printArea(triangle)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.base * s.height
}
