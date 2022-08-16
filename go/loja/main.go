package main

import (
	"log"
	"loja/routes"
	"net/http"
)

func main() {
	router := routes.CarregaRotas()
	log.Fatal(http.ListenAndServe(":8001", router))
}
