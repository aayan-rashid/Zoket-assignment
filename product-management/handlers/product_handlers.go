package handlers

import (
	"encoding/json"
	"net/http"
	"product-management/models"
	"strconv"

	"github.com/gorilla/mux"
)

var products []models.Product

// CreateProduct handles POST /products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	product.ID = len(products) + 1
	products = append(products, product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProductByID handles GET /products/{id}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// GetProducts handles GET /products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	// priceRange := r.URL.Query().Get("price_range")
	name := r.URL.Query().Get("name")

	var filteredProducts []models.Product
	for _, product := range products {
		if (userID == "" || strconv.Itoa(product.UserID) == userID) &&
			(name == "" || product.ProductName == name) {
			filteredProducts = append(filteredProducts, product)
		}
	}
	json.NewEncoder(w).Encode(filteredProducts)
}
