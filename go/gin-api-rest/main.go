package main

import (
	"gin-api-rest/src/database"
	"gin-api-rest/src/routes"
)

func init() {
	database.ConectaComBancoDeDados()
}

func main() {
	routes.HandleRequests()
}
