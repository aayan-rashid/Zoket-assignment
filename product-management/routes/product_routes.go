package routes

import (
	"product-management/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
}
