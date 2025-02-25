package main

import "fmt"

func main() {
	fmt.Println("Learn If else Condition in Go")
	// if else condition
	// var age int = 10
	age := 10
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else if age < 18 {
		fmt.Println("You are not eligible to vote")
	} else {
		fmt.Println("Invalid age")
	}
}
