package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	searchProducts := models.SearchProducts()
	templates.ExecuteTemplate(w, "Index", searchProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	createProducts := models.CreateNewProduct
	templates.ExecuteTemplate(w, "New", createProducts)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")
		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}
		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeletProduct(productId)
	http.Redirect(w, r, "/", 301)
}
