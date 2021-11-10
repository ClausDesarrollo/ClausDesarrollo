/**
 * main.go
 * Golang webserver with concurrency
 *
 */

package main

import (
	"log"
	"net/http"

	output "go-webserver/views"
)

func main() {

	http.HandleFunc("/calculate", output.ShapeResponse)
	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
