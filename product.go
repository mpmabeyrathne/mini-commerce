package main

import (
	"errors"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
	ID    int    `json:"id"`
}

type InStockProducts struct {
	Name   string
	Stock  int
	Status string
}

func (p *Product) IsAvailable() bool {
	return p.Stock > 0
}

func reduceStockByOne(stockPointer *int) int {
	*stockPointer = *stockPointer - 1
	return *stockPointer

}

func addProduct(product Product, products []Product) []Product {
	return append(products, product)
}

func findProductByName(value string, products []Product) (Product, error) {
	for _, p := range products {
		if p.Name == value {
			return p, nil
		}
	}
	return Product{}, errors.New("Not found")
}

func findProductById(value int, products []Product) (Product, error) {
	for _, p := range products {
		if p.ID == value {
			return p, nil
		}
	}
	return Product{}, errors.New("Not found")
}

func calculateTotal(price int, quantity int) int {
	return price * quantity
}

func reduceStock(productName string, quantity int, products []Product) []Product {
	for i := range products {
		if products[i].Name == productName && (products[i].Stock >= quantity) {
			products[i].Stock = products[i].Stock - quantity
		}
	}

	return products

}

func getStockStatus(products []Product) []InStockProducts {
	var inStockProducts []InStockProducts
	for _, p := range products {
		var status string
		switch {
		case p.Stock == 0:
			status = "out of stock"
		case p.Stock > 0 && p.Stock < 4:
			status = "low stock"
		case p.Stock >= 4:
			status = "in stock"
		default:
			status = "invalid stock"
		}
		inStockProducts = append(inStockProducts, InStockProducts{
			Name:   p.Name,
			Stock:  p.Stock,
			Status: status,
		})

	}

	return inStockProducts
}

var products = []Product{
	{ID: 1, Name: "Mechanical Keyboard", Price: 8500, Stock: 5},
	{ID: 2, Name: "Gaming Mouse", Price: 4500, Stock: 8},
	{ID: 3, Name: "USB-C Hub", Price: 6500, Stock: 3},
}
