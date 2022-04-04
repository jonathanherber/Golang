package main

import (
	"fmt"

	"github.com/jonathanherber/Golang_api/models"
	"github.com/jonathanherber/Golang_api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{ID: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor...")
	routes.HandleRequest()
}
