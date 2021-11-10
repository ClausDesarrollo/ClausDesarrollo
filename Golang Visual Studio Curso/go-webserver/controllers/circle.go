package controllers

import (
	"fmt"

	"go-webserver/models/structs"
)

func Circle(channel chan string, response *[]structs.Output) {
	circle := structs.Circle{Diameter: 20}
	circleArea := circle.Area()
	circlePerimeter := circle.Perimeter()

	circleResponse := structs.Output{
		Shape:   "Círculo",
		Message: fmt.Sprintf("Área: %.2f  Perímetro: %.2f", circleArea, circlePerimeter),
	}

	*response = append(*response, circleResponse)

	channel <- fmt.Sprintf(".: Círculo :.\nÁrea: %.2f\tPerímetro: %.2f\n", circleArea, circlePerimeter)
}
