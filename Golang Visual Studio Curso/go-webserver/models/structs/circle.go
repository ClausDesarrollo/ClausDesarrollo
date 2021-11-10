package structs

import "math"

type Circle struct {
	Diameter float32
}

func (circle Circle) Area() float64 {
	return 3.1416 * math.Pow(float64(circle.Diameter/2), 2)
}

func (circle Circle) Perimeter() float32 {
	return 2 * 3.1416 * (circle.Diameter / 2)
}
