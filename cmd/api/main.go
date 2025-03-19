package main

import (
	"fmt"
	"net/http"

	"github.com/tfrancar/pizzaria/internal/router"
)

func main() {
	// repository.LoadPizzas()
	fmt.Println("Starting the Pizzeria server...")

	// Register the routes
	router := router.PizzasRouters()

	// Server up on port 8080
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
