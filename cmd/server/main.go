package main

import (
	"fmt"
	"net/http"
	"github.com/72sevenzy2/golang-API/internal/handler"
	"github.com/72sevenzy2/golang-API/internal/health"
	"github.com/72sevenzy2/golang-API/internal/service"
)

func main() {
	service := &service.GreetCounter{}

	http.HandleFunc("/greet", handler.GreetHandler(service))
	http.HandleFunc("/health", health.HealthHandler())

	fmt.Println("API running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
