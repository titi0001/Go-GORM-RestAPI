package main

import (
	"fmt"

	"github.com/titi0001/Go-GORM-RestAPI/database"
	"github.com/titi0001/Go-GORM-RestAPI/routes"
)

func main() {

	fmt.Println("Inciando o servidor")
	database.ConectaDb()
	routes.HandleRequest()
}
