package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// create struct for product
type Rating struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}
type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

func main() {
	fmt.Println("We are learning CRUD in Go")
	// CRUD - Create, Read, Update, Delete
	// Create
	res, err := http.Get("https://fakestoreapi.com/products/3")

	if err != nil {
		fmt.Println("Error while fetching data", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while fetching data:", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error while reading data", err)
	// 	return
	// }
	// fmt.Println(string(data))
	var product Product
	err = json.NewDecoder(res.Body).Decode(&product)
	if err != nil {
		fmt.Println("Error while decoding data", err)
		return
	}
	// fmt.Println(product)
	fmt.Println("Product ID: ", product.Id)
	fmt.Println("Product Title: ", product.Title)
	fmt.Println("Product Price: ", product.Price)
	// fmt.Println("Product Description: ", product.Description)
	fmt.Println("Product Category: ", product.Category)
	fmt.Println("Product Image: ", product.Image)
	fmt.Println("Product Rating: ", product.Rating)
	fmt.Println("Product Rating Rate: ", product.Rating.Rate)
	fmt.Println("Product Rating Count: ", product.Rating.Count)

}
