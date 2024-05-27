package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}
type Triangle struct {
	base   float64
	height float64
}
type Circle struct {
	radius float64
}
type Shape interface {
	Area() float64
}

// func Perimeter(rec Rectangle) float64 {
// 	return 2 * (rec.width + rec.height)
// }

// func Area(rec Rectangle) float64 {
// 	return rec.width * rec.height
// }

func (r Rectangle) Area() float64 {
	return r.width * r.height

}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius

}
func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}
