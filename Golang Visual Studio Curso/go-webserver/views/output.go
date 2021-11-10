package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	shape "go-webserver/controllers"
	"go-webserver/models/structs"
)

func ShapeResponse(w http.ResponseWriter, r *http.Request) {
	output := make([]structs.Output, 0)
	channel := make(chan string)

	go shape.Square(channel, &output)
	go shape.Rectangule(channel, &output)
	go shape.Triangule(channel, &output)
	go shape.Circle(channel, &output)

	for i := 0; i < 4; i++ {
		fmt.Print(<-channel)
	}

	close(channel)

	time.Sleep(1 * time.Second)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
