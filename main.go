package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {

	const appName string = "Mini Commerce"
	cfg, err := LoadConfig()
	if err != nil {
		log.Println("Invalid port configuration:", err)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", HealthHandler)

	mux.HandleFunc("GET /products", GetProductstHandler)

	mux.HandleFunc("GET /product/{id}", GetProducByIdtHandler)

	mux.HandleFunc("POST /product", CreateProductHandler)

	mux.HandleFunc("DELETE /product/{id}", DeleteProductHandler)

	mux.HandleFunc("PUT /product/{id}", UpdateProductHandler)

	log.Printf("Starting %s %s server on :%d", appName, cfg.Env, cfg.Port)

	err = http.ListenAndServe(":"+strconv.Itoa(cfg.Port), mux)

	if err != nil {
		log.Fatal("Server error:", err)
	}
}
