package main

import "fmt"

func main() {
	fmt.Println("Started Error Handling in Go")
	// ans, _ := div(10, 0)
	ans, err := div(10, 0)
	if err != nil {
		fmt.Println("Division is ", ans)
	}
}
func div(x, y float32) (float32, error) {
	if y == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return x / y, nil
}
