package _5_structs_interfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Areable interface {
	Area() float64
}

type Perimeterable interface {
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

type Answer struct {
}

func (a Answer) Perimeter() float64 {
	return 42
}

func (a Answer) Area() float64 {
	return 42
}
