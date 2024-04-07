package routes

import (
	"log"
	"net/http"

	"github.com/titi0001/Go-GORM-RestAPI/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
