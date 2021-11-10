package structs

type Triangule struct {
	Base   float32
	Height float32
	SizeA  float32
	SizeB  float32
}

func (triangule Triangule) Area() float32 {
	return (triangule.Base * triangule.Height) / 2
}

func (triangule Triangule) Perimeter() float32 {
	return triangule.Base + triangule.SizeA + triangule.SizeB
}
