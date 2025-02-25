package main

import (
	"fmt"
)

func main() {
	var name string = "Bharti"
	fmt.Println(name)
	var version = 76
	fmt.Println(version)
	var curruncy int = 100
	fmt.Println(curruncy)
	curruncy = 200
	fmt.Println(curruncy)
	const pi = 3.14
	fmt.Println(pi)
	// pi = 3.15 // cannot assign to pi
	var PublicVarriable = "Public varriable"
	fmt.Println(PublicVarriable)
	var privateVarriable = "Private varriable"
	fmt.Println(privateVarriable)
	// in go lang we can not use private varriable outside the package but we can use public varriable outside the package
	// varriable name should be start with small letter if we want to make it private
	// varriable name should be start with capital letter if we want to make it public
}
