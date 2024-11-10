package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/FabioLeitao/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPesonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/id/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
