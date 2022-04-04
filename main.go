package main

import (
	"fmt"

	"github.com/jonathanherber/Golang_api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor...")
	routes.HandleRequest()
}
