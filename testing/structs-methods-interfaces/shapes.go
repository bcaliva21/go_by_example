package main

import "math"

type Shape interface {
	Area() float64
}

type Rect struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {}
