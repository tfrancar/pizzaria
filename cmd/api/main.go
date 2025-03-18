package main

import (
	"fmt"
	"net/http"

	"github.com/tfrancar/pizzaria/internal/router"
)

func main() {
	fmt.Println("Iniciando servidor da Pizzaria...")

	// Criar roteador
	r := router.PizzasRouters()

	// Subir servidor na porta 8080
	http.ListenAndServe(":8080", r)
}
