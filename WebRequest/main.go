package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("We are learning Web Request in Go")
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer response.Body.Close()
	// fmt.Println(response)
	fmt.Printf("type of response is %T\n", response)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Println(string(data))
	fmt.Printf("type of data is %T\n", data)

}
