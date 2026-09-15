package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	maxID := 0

	var newProd Product
	err := json.NewDecoder(r.Body).Decode(&newProd)

	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if newProd.Name == "" {
		http.Error(w, "Product name is required", http.StatusBadRequest)
		return
	}

	if newProd.Price <= 0 {
		http.Error(w, "Price must be greater than 0", http.StatusBadRequest)
		return
	}

	if newProd.Stock < 0 {
		http.Error(w, "Stock cannot be negative", http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	newProd.ID = maxID + 1

	products = append(products, newProd)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProd)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedData Product
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if updatedData.Name == "" {
		http.Error(w, "Product name is required", http.StatusBadRequest)
		return
	}
	if updatedData.Price <= 0 {
		http.Error(w, "Price must be greater than 0", http.StatusBadRequest)
		return
	}
	if updatedData.Stock < 0 {
		http.Error(w, "Stock cannot be negative", http.StatusBadRequest)
		return
	}

	foundIndex := -1
	for i, p := range products {
		if p.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	updatedData.ID = id
	products[foundIndex] = updatedData

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedData)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID layout", http.StatusBadRequest)
		return
	}
	foundIndex := -1
	for i, p := range products {
		if p.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	products = append(products[:foundIndex], products[foundIndex+1:]...)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product with ID %d removed successfully", id)
}

func GetProducByIdtHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID layout", http.StatusBadRequest)
		return
	}
	product, err := findProductById(id, products)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProductstHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}
