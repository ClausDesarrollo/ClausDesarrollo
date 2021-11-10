package controllers

import (
	"fmt"

	"go-webserver/models/structs"
)

func Triangule(channel chan string, response *[]structs.Output) {
	triangule := structs.Triangule{Base: 12, Height: 15, SizeA: 16.6, SizeB: 16.6}
	trianguleArea := triangule.Area()
	triangulePerimeter := triangule.Perimeter()

	circleResponse := structs.Output{
		Shape:   "Triángulo",
		Message: fmt.Sprintf("Área: %.2f  Perímetro: %.2f", trianguleArea, triangulePerimeter),
	}

	*response = append(*response, circleResponse)

	channel <- fmt.Sprintf(".: Triángulo :.\nÁrea: %.2f\tPerímetro: %.2f\n", trianguleArea, triangulePerimeter)
}
