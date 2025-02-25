package main

import "fmt"

func main() {
	fmt.Println("Learn for loop in Go")
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// for infinite loop
	// counter := 0
	// for {
	// 	fmt.Println(counter)
	// 	counter++
	// 	if counter == 5 {
	// 		break
	// 	}
	// counter := 10
	// for {
	// 	fmt.Println(counter)
	// 	counter--
	// 	if counter == 5 {
	// 		continue
	// 	}
	// 	if counter == 0 {
	// 		break
	// 	}
	// }
	// for loop with range
	var numbers = []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)

	}
}
