package main

import (
	"encoding/json"
	"fmt"
)

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
	fmt.Println(newProduct.IsAvailable())
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

	searchedProduct, err := findProductByName("Gaming Mouse", products)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(
		"\nFound: %s\nPrice: %d\nStock: %d\n",
		searchedProduct.Name,
		searchedProduct.Price,
		searchedProduct.Stock,
	)

	newProducts := reduceStock("Gaming Mouse", 2, products)

	fmt.Println(newProducts)

	productStock := getStockStatus(products)
	fmt.Println(productStock)

	categories := map[string]string{
		"keyboard": "Accessories",
		"mouse":    "Accessories",
		"laptop":   "Computers",
	}

	fmt.Printf(categories["keyboard"])

	stock := 5
	fmt.Println(reduceStockByOne(&stock))
	fmt.Println(stock)

	jsonBytes, err := json.Marshal(newProduct)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
