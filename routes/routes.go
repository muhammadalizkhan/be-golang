package routes

import (
	"github.com/EENmalikhanDev/be-golang/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", handlers.GetProductsHandler).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProductHandler).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProductHandler).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProductHandler).Methods("DELETE")

	return router
}
