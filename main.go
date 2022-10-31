package main

import (
	"log"

	"github.com/murilorscampos/gin-api-rest/database"
	"github.com/murilorscampos/gin-api-rest/routes"
)

func main() {

	log.Println("Conectando ao banco de dados...")
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
