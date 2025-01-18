package main

import (
	"log"
	"net/http"
	"product-api/routes"
)

func main() {
	router := routes.RegisterRoutes()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
