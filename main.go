package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	const appName string = "Mini Commerce"
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "development"
	}
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port configuration:", err)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthHandler)

	mux.HandleFunc("GET /products", GetProductstHandler)

	mux.HandleFunc("GET /product/{id}", GetProducByIdtHandler)

	mux.HandleFunc("POST /product", CreateProductHandler)

	mux.HandleFunc("DELETE /product/{id}", DeleteProductHandler)

	mux.HandleFunc("PUT /product/{id}", UpdateProductHandler)

	log.Printf("Starting %s %s server on :%d", appName, environment, port)

	err = http.ListenAndServe(":"+strconv.Itoa(port), mux)

	if err != nil {
		log.Fatal("Server error:", err)
	}
}
