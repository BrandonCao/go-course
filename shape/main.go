package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func main() {
	triangle := Triangle{base: 10, height: 10}
	square := Square{sideLength: 10}

	fmt.Println("Triangle Area:", triangle.getArea())
	fmt.Println("Square Area:", square.getArea())
}
