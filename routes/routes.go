package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/FabioLeitao/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPesonalidades)
	log.Fatal(http.ListenAndServe(":8000", r))
}
