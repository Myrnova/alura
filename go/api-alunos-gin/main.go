package main

import (
	"api-alunos-gin/src/database"
	"api-alunos-gin/src/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
