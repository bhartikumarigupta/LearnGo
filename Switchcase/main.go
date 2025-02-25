package main

import "fmt"

func main() {
	fmt.Println("Learn Switch Case in Go")
	// switch case
	var age int = 10
	switch age {
	case 10:
		fmt.Println("You are 10 years old")
	case 20:
		fmt.Println("You are 20 years old")
	case 30:
		fmt.Println("You are 30 years old")
	default:
		fmt.Println("Invalid age")
	}
}
