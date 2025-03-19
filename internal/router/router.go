package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tfrancar/pizzaria/internal/handlers"
	"github.com/tfrancar/pizzaria/internal/middleware"
)

func PizzasRouters() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.LoggerMiddleware)

	// Root route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Golang's Pizzaria API!")
	})

	// Pizza routes
	r.Get("/pizzas", handlers.ListPizzas)
	r.Post("/pizzas", handlers.CreatePizza)
	r.Get("/pizzas/{id}", handlers.GetPizza)
	r.Put("/pizzas/{id}", handlers.UpdatePizza)
	r.Delete("/pizzas/{id}", handlers.DeletePizza)

	return r
}
