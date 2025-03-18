package repository

import (
	"time"

	"github.com/tfrancar/pizzaria/internal/models"
)

var pizzas = []models.Pizza{
	{ID: 1, Name: "Margherita", Description: "Classic pizza with tomatoes, mozzarella, basil, pasta sauce and olive oil.", Price: 52.99, Size: "large"},
	{ID: 2, Name: "Lombo com catupiry", Description: "Pizza with salami, mozzarella, olives, garlic and catupiry cheese.", Price: 14.99, Size: "medium"},
	{ID: 3, Name: "calabrese sausage with onions", Description: "Pizza with Italian calabrese sausage, mozzarella, olives, garlic and onions rings.", Price: 4.99, Size: "small"},
	{ID: 4, Name: "four cheese pizza", Description: "Pizza with gorgonzola ,provolone, parmesan, mozzarella, olives and pasta suace.", Price: 54.99, Size: "extra large"},
}

// GetPizzas returns all pizzas
func GetPizzas() []models.Pizza {
	return pizzas
}

// CreatePizza creates a new pizza
func AddPizza(pizza models.Pizza) (*models.Pizza, error) {

	newPizza := &models.Pizza{
		Name:        pizza.Name,
		Description: pizza.Description,
		Price:       pizza.Price,
		Size:        pizza.Size,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, *newPizza)
	return newPizza, nil
}
