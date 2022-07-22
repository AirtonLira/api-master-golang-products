package controllers

import (
	"api-master-golang-products/database"
	"api-master-golang-products/entities"
	"api-master-golang-products/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if utils.ValidJWT(w, r) {
		w.Header().Set("Content-Type", "application/json")
		var product entities.Product
		json.NewDecoder(r.Body).Decode(&product)
		database.Instance.Create(&product)
		json.NewEncoder(w).Encode(product)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not authorized to create product"))
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	if utils.ValidJWT(w, r) {
		productId := mux.Vars(r)["id"]
		if !checkIfProductExists(productId) {
			json.NewEncoder(w).Encode("Product not found")
			return
		}
		var product entities.Product
		database.Instance.First(&product, productId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not authorized to get product"))
	}
}

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Instance.First(&product, productId)

	return product.ID != 0
}
