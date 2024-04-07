package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/titi0001/Go-GORM-RestAPI/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", r))

}
