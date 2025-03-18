package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tfrancar/pizzaria/internal/repository"
)

// List all pizzas
func ListPizzas(w http.ResponseWriter, r *http.Request) {
	pizzas := repository.GetPizzas()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizzas)
	w.WriteHeader(http.StatusOK)
}
