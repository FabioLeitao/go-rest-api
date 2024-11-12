package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/FabioLeitao/go-rest-api/controllers"
	"github.com/FabioLeitao/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPesonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/id/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/id/{id}", controllers.ApagaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/id/{id}", controllers.EditaUmaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
