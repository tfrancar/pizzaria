package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tfrancar/pizzaria/internal/models"
	"github.com/tfrancar/pizzaria/internal/repository"
)

// List all pizzas
func ListPizzas(w http.ResponseWriter, r *http.Request) {
	pizzas := repository.GetPizzas()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizzas)
	w.WriteHeader(http.StatusOK)
}

// Create a new pizza
func CreatePizza(w http.ResponseWriter, r *http.Request) {
	var newPizza models.Pizza

	// CreatePizza add a new pizza
	if err := json.NewDecoder(r.Body).Decode(&newPizza); err != nil {
		http.Error(w, "Something went wrong decoding. Try again later", http.StatusBadRequest)
		return
	}

	// Create a new pizza object
	pizza, err := repository.AddPizza(newPizza)
	if err != nil {
		http.Error(w, "Failed to create pizza", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizza)
	w.WriteHeader(http.StatusCreated)
}
