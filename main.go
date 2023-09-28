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
	}

	fmt.Println("Iniciando o servidor REST com Go")
	routes.HandleRequest()

}
