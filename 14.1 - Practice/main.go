package main

import "fmt"

type Product struct {
	title    string
	id       int
	price    float64
	products []string
}

type Product2 struct {
	title string
	id    int
	price float64
}

func main() {
	// array within structure
	new := Product{
		title:    "books",
		id:       45,
		price:    4556.3,
		products: []string{"charlie", "goat"},
	}

	fmt.Println(new)

	new.products = append(new.products, "cat")

	fmt.Println(new)

	// array of structures
	products := []Product2{
		{
			"animals",
			43,
			500.1,
		},
		{
			"fish",
			23,
			34.32,
		},
	}
	fmt.Println(products)

	// update products list
	newProduct := Product2{
		"sports",
		50,
		603.32,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}
