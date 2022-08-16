package controllers

import (
	"html/template"
	"log"
	"loja/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

	if err != nil {
		log.Println("Erro na conversão do preço: ", err)
	}

	quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

	if err != nil {
		log.Println("Erro na conversão da quantidade: ", err)
	}

	models.CriarNovoProduto(models.Produto{Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade})

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	models.DeletaProduto(id)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	produto := models.BuscaProduto(id)

	templates.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Erro na conversão do id: ", err)
	}
	nome := r.FormValue("nome")
	descricao := r.FormValue("descricao")
	preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

	if err != nil {
		log.Println("Erro na conversão do preço: ", err)
	}

	quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

	if err != nil {
		log.Println("Erro na conversão da quantidade: ", err)
	}

	models.EditaProduto(models.Produto{Id: id, Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade})

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
