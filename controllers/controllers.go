package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/titi0001/Go-GORM-RestAPI/database"
	"github.com/titi0001/Go-GORM-RestAPI/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
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

func CriaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)

}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Personalidade

	database.DB.Delete(&person, id)
	json.NewEncoder(w).Encode(person)

}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Personalidade

	database.DB.First(&person, id)
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Save(&person)
	json.NewEncoder(w).Encode(person)
}
