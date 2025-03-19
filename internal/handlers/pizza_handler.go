package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/tfrancar/pizzaria/internal/models"
	"github.com/tfrancar/pizzaria/internal/repository"
)

// List all pizzas
func ListPizzas(w http.ResponseWriter, r *http.Request) {
	pizzas := repository.GetPizzas()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizzas)
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

}

// Get a pizza by ID
func GetPizza(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pizza, found := repository.GetPizzaByID(id)
	if !found {
		http.Error(w, "Pizza not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizza)

}

// Update a pizza by ID
func UpdatePizza(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedPizza models.Pizza

	// Decode the JSON request body into updatedPizza
	if err := json.NewDecoder(r.Body).Decode(&updatedPizza); err != nil {
		http.Error(w, "Something went wrong decoding. Try again later", http.StatusBadRequest)
		return
	}

	// Update the pizza in the repository
	pizza, found := repository.UpdatePizza(id, updatedPizza)
	if !found {
		http.Error(w, "Pizza not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pizza)
}

// Delete a pizza by ID
func DeletePizza(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	deleted := repository.DeletePizza(id)
	if !deleted {
		http.Error(w, "Pizza not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
