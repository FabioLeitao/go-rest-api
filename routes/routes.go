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
	r.HandleFunc("/api/personalidades", controllers.TodasPesonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/id/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/id/{id}", controllers.ApagaUmaPersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}
