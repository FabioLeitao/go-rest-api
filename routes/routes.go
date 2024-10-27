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
	r.HandleFunc("/api/personalidades", controllers.TodasPesonalidades)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade)
	log.Fatal(http.ListenAndServe(":8000", r))
}
