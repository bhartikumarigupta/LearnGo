package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello", name, "let's start the journey")
	// fmt.scan reads untill the first wwitespace character so if we want to read the whole line we can use bufio package newReader method or newscanner method
	fmt.Println("Please enter your full name")
	reader := bufio.NewReader(os.Stdin)
	fullname, _ := reader.ReadString('\n')
	fmt.Println("Hello", fullname, "let's start the journey")
}
