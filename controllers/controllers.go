package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/FabioLeitao/go-rest-api/database"
	"github.com/FabioLeitao/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page!")
}

func TodasPesonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Personalidade
	database.DB.First(&person, id)
	json.NewEncoder(w).Encode(person)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPessoa models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPessoa)
	database.DB.Create(&novaPessoa)
	json.NewEncoder(w).Encode(novaPessoa)
}
