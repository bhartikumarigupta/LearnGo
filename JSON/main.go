package main

import (
	"encoding/json"
	"fmt"
)

// person name ka structure
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Printf("Learn JSON in Go\n")
	person := Person{
		Name:    "Raj",
		Age:     25,
		City:    "Delhi",
		IsAdult: true,
	}
	fmt.Println(person)
	//convert struct (person) to json
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}
	fmt.Println("Person data is ", string(jsonData))
	// convert json to struct (person) decodig (unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling: ", err)
		return
	}
	fmt.Println("Person data is ", personData)
	fmt.Println("Name: ", personData.Name)
	fmt.Println("Age: ", personData.Age)
	fmt.Println("City: ", personData.City)
	fmt.Println("IsAdult: ", personData.IsAdult)

}
