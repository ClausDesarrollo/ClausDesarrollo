package controllers

import (
	"fmt"

	"go-webserver/models/structs"
)

func Square(channel chan string, response *[]structs.Output) {
	square := structs.Square{Width: 5}
	squareArea := square.Area()
	squarePerimeter := square.Perimeter()

	circleResponse := structs.Output{
		Shape:   "Cuadrado",
		Message: fmt.Sprintf("Área: %d  Perímetro: %d", squareArea, squarePerimeter),
	}

	*response = append(*response, circleResponse)

	channel <- fmt.Sprintf(".: Cuadrado :.\nÁrea: %d\tPerímetro: %d\n", squareArea, squarePerimeter)
}
