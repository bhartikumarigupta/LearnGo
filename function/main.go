package main

import "fmt"

func main() {
	fmt.Println("Learn functions in Go")
	// function call
	greet()
	// function call with return value
	fmt.Println(Add(10, 20))
	fmt.Println(Sub(20, 10))
	fmt.Println(Mul(20, 10))
	fmt.Println(Div(20, 10))
	fmt.Println(Calculator(20, 10))
	fmt.Println(Add1(20, 10))
}
func greet() {
	fmt.Println("Good Morning")
}
func Add(x int, y int) int {
	return x + y
}
func Sub(x, y int) int {
	return x - y
}
func Mul(x, y int) int {
	return x * y
}
func Div(x, y int) int {
	return x / y
}
func Calculator(x, y int) (int, int, int, int) {
	return x + y, x - y, x * y, x / y
}
func Add1(x, y int) (result int) {
	result = x + y
	return
}
