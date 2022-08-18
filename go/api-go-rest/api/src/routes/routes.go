package routes

import (
	"api-go-rest/src/controllers"
	"api-go-rest/src/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades", controllers.TodasAsPersonalidades).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades/{id}", controllers.BuscaUmaPersonalidade).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades", controllers.CriaUmNovaPersonalidade).Methods(http.MethodPost)
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods(http.MethodDelete)
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaUmaPersonalidade).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
