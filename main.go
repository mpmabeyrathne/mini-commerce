package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
	Stock int
}

type InStockProducts struct {
	Name   string
	Stock  int
	Status string
}

func main() {
	const appName string = "Mini Commerce"
	port := 8080
	var environment string = "development"

	fmt.Printf("Application: %s \n", appName)
	fmt.Printf("Port: %d \n", port)
	fmt.Printf("Environment: %s\n", environment)

	products := []Product{{
		Name:  "Mechanical Keyboard",
		Price: 8500,
		Stock: 5,
	}, {
		Name:  "Gaming Mouse",
		Price: 4500,
		Stock: 8,
	}, {
		Name:  "USB-C Hub",
		Price: 6500,
		Stock: 3,
	},
	}

	quantity := 2

	total := calculateTotal(products[0].Price, quantity)

	fmt.Printf("Product: %s\n", products[0].Name)
	fmt.Printf("Price: %d\n", products[0].Price)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Total: %d\n", total)

	if quantity <= products[0].Stock {
		fmt.Printf("%s is available\n", products[0].Name)
	} else {
		fmt.Printf("%s is out of stock\n", products[0].Name)
	}

	newProduct := Product{
		Name:  "Laptop Stand",
		Price: 7000,
		Stock: 4,
	}

	products = addProduct(newProduct, products)

	// fmt.Printf("Total products: %d\n", len(products))

	// for x := 0; x < len(products); x++ {
	// 	fmt.Printf("\nProduct: %s\n", products[x].Name)
	// 	fmt.Printf("Price: %d\n", products[x].Price)
	// 	fmt.Printf("Stock: %d\n", products[x].Stock)

	// }

	// for i := range len(products) {
	// 	fmt.Printf("\nProduct: %s\n", products[i].Name)
	// 	fmt.Printf("Price: %d\n", products[i].Price)
	// 	fmt.Printf("Stock: %d\n", products[i].Stock)
	// }

	for _, product := range products {
		fmt.Printf("\nProduct: %s\n", product.Name)
		fmt.Printf("Price: %d\n", product.Price)
		fmt.Printf("Stock: %d\n", product.Stock)
	}

	searchedProduct, productStatus := findProductByName("Gaming Mouse", products)

	if productStatus {
		fmt.Printf("\nFound: %s\nPrice: %d\nStock: %d\n", searchedProduct.Name, searchedProduct.Price, searchedProduct.Stock)
	} else {
		fmt.Printf("\nProduct not found")
	}

	newProducts := reduceStock("Gaming Mouse", 2, products)

	fmt.Println(newProducts)

	productStock := getStockStatus(products)
	fmt.Println(productStock)
}

func addProduct(product Product, products []Product) []Product {
	return append(products, product)
}

func findProductByName(value string, products []Product) (Product, bool) {
	for _, p := range products {
		if p.Name == value {
			return p, true
		}
	}
	return Product{}, false
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
