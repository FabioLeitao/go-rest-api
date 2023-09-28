package main

import (
	"fmt"

	"github.com/FabioLeitao/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor REST com Go")
	routes.HandleRequest()

}
