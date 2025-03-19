package repository

import (
	"time"

	"github.com/tfrancar/pizzaria/internal/models"
)

// var pizzas = []models.Pizza{
// 	{ID: 1, Name: "Margherita", Description: "Classic pizza with tomatoes, mozzarella, basil, pasta sauce and olive oil.", Price: 52.99, Size: "large"},
// 	{ID: 2, Name: "Lombo com catupiry", Description: "Pizza with salami, mozzarella, olives, garlic and catupiry cheese.", Price: 14.99, Size: "medium"},
// 	{ID: 3, Name: "calabrese sausage with onions", Description: "Pizza with Italian calabrese sausage, mozzarella, olives, garlic and onions rings.", Price: 4.99, Size: "small"},
// 	{ID: 4, Name: "four cheese pizza", Description: "Pizza with gorgonzola ,provolone, parmesan, mozzarella, olives and pasta suace.", Price: 54.99, Size: "extra large"},
// }

var pizzas []models.Pizza

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
	// savePizzas()
	return newPizza, nil
}

// GetPizzaByID retrieves a pizza by its ID
func GetPizzaByID(id int) (*models.Pizza, bool) {
	for _, pizza := range pizzas {
		if pizza.ID == id {
			return &pizza, true
		}
	}
	return nil, false
}

// UpdatePizza updates an existing pizza
func UpdatePizza(id int, updatedPizza models.Pizza) (*models.Pizza, bool) {
	for i, pizza := range pizzas {
		if pizza.ID == id {
			pizzas[i].Name = updatedPizza.Name
			pizzas[i].Description = updatedPizza.Description
			pizzas[i].Price = updatedPizza.Price
			pizzas[i].Size = updatedPizza.Size
			pizzas[i].UpdatedAt = time.Now()
			// savePizzas()
			return &pizzas[i], true
		}
	}
	return nil, false
}

// DeletePizza deletes a pizza by its ID
func DeletePizza(id int) bool {
	for i, pizza := range pizzas {
		if pizza.ID == id {
			// Descompactando o slice. Os 3 pontos significa(unpacking) ou spread operator
			// Unpacking the slice. The three dots mean unpacking or spread operator.
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			// savePizzas()
			return true
		}
	}
	return false
}

// // Simulating loading pizzas from a database
// func LoadPizzas() []models.Pizza {
// 	file, err := os.Open("data/pizza.json")
// 	if err != nil {
// 		fmt.Println("Error open file: ", err)
// 		return []models.Pizza{}
// 	}
// 	defer file.Close()

// 	decode := json.NewDecoder(file)
// 	if err = decode.Decode(&pizzas); err != nil {
// 		fmt.Println("Error decoding JSON: ", err)
// 		return []models.Pizza{}
// 	}

// 	return pizzas

// }

// Simulating saving pizzas to a database
// func savePizzas() {
// 	file, err := os.Create("data/pizza.json")
// 	if err != nil {
// 		fmt.Println("Error creating file: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	if err = encoder.Encode(pizzas); err != nil {
// 		fmt.Println("Error encoding JSON: ", err)
// 		return
// 	}

// }
