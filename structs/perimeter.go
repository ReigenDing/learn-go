package structs

import (
	"math"
)

func Perimeter(rectangle RecTangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2

}

func Area(rectangle RecTangle) float64 {
	return rectangle.Width * rectangle.Height
}

type RecTangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// 会报重复定义错误
// func Area(circle Circle) float64 {
// 	return 0.0
// }

func (c Circle) Area() float64 {
	// fmt.Printf("reciver %v  address %p, parameter %v address %p", c, &c, circle, &circle)
	return math.Pi * c.Radius * c.Radius
}

func (r RecTangle) Area() float64 {
	return r.Height * r.Width
}

type Shape interface {
	Area() float64
}
