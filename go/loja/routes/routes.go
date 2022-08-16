package routes

import (
	"loja/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func CarregaRotas() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index).Methods(http.MethodGet)
	router.HandleFunc("/new", controllers.New).Methods(http.MethodGet)
	router.HandleFunc("/insert", controllers.Insert).Methods(http.MethodPost)
	router.HandleFunc("/delete/{id}", controllers.Delete).Methods(http.MethodGet)
	router.HandleFunc("/edit/{id}", controllers.Edit).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", controllers.Update).Methods(http.MethodPost)


	return router
}
