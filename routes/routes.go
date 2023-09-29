package routes

import (
	"log"
	"net/http"

	"github.com/FabioLeitao/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPesonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
