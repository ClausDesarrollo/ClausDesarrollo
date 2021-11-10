package structs

type Square struct {
	Width int
}

func (square Square) Area() int {
	return square.Width * square.Width
}

func (square Square) Perimeter() int {
	return square.Width * 4
}
