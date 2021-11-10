package structs

type Rectangule struct {
	Width  int
	Height int
}

func (rectangule Rectangule) Area() int {
	return rectangule.Width * rectangule.Height
}

func (rectangule Rectangule) Perimeter() int {
	return (rectangule.Width * 2) + (rectangule.Height * 2)
}
