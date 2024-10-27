package main

import (
	"fmt"

	"github.com/FabioLeitao/go-rest-api/models"
	"github.com/FabioLeitao/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
		{Nome: "Nome 3", Historia: "Historia 3"},
		{Nome: "Nome 4", Historia: "Historia 4"},
	}

	fmt.Println("Iniciando o servidor REST com Go")
	routes.HandleRequest()

}
