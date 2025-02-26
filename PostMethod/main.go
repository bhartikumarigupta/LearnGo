package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

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
	fmt.Println("We are learning Post Method in Go")

	product := Product{
		Id:          201,
		Title:       "Women Short Kurti",
		Price:       200,
		Description: "This is a short women cotton kurti with beautiful design and color options",
		Category:    "Women`s Clothing",
		Image:       "https://fakestoreapi.com/img/71z3kpMAYsL._AC_UY879_.jpg",
		Rating: Rating{
			Rate:  4.5,
			Count: 100,
		},
	}
	// convert product to json
	jsonData, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}
	//convert json data to string
	jsonString := string(jsonData)
	// url of fakestoreapi server
	myurl := "https://fakestoreapi.com/products"
	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	// post method to create a new product on fakestoreapi server using post method and json data as body of request
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error while posting data", err)
		return
	}
	defer res.Body.Close()
	// if res.StatusCode != http.StatusCreated {
	// 	fmt.Println("Error while posting data:", res.Status)
	// 	return
	// }
	fmt.Println("Product created successfully")
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading data", err)
		return
	}
	fmt.Println("Response: ", string(data))

}
