package controllers

import (
	"fmt"

	"go-webserver/models/structs"
)

func Rectangule(channel chan string, response *[]structs.Output) {
	rectangule := structs.Rectangule{Width: 5, Height: 8}
	rectanguleArea := rectangule.Area()
	rectangulePerimeter := rectangule.Perimeter()

	circleResponse := structs.Output{
		Shape:   "Rectángulo",
		Message: fmt.Sprintf("Área: %d  Perímetro: %d", rectanguleArea, rectangulePerimeter),
	}

	*response = append(*response, circleResponse)

	channel <- fmt.Sprintf(".: Rectángulo :.\nÁrea: %d\tPerímetro: %d\n", rectanguleArea, rectangulePerimeter)
}
