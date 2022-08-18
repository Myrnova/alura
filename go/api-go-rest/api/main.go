package main

import (
	"api-go-rest/src/database"
	"api-go-rest/src/routes"
	"fmt"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com GO!")
	routes.HandleRequest()
}
