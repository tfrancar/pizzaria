package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PizzasRouters() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Golang's Pizzaria API!")
	})

	// Pizza routes
	// r.Get("/pizzas", getPizzas)
	// r.Post("/pizzas", createPizza)
	// r.Get("/pizzas/{id}", getPizza)
	// r.Put("/pizzas/{id}", updatePizza)
	// r.Delete("/pizzas/{id}", deletePizza)

	return r
}
