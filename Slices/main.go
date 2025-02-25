package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	fmt.Println(number)
	number = append(number, 6, 7, 8)
	fmt.Println(number)
	name := []string{}
	fmt.Println(name)
	name = append(name, "Bharti", "Raj", "Rahul", "Ravi", "Rohit")
	fmt.Printf("%q", name)
	// slices through make function
	Number := make([]int, 5, 10) // 5 is length and 10 is capacity
	fmt.Println(Number)
	Number = append(Number, 1, 2, 3, 4, 5)
	fmt.Println(Number)
	fmt.Println(len(Number))
	fmt.Println(cap(Number))
	Number = append(Number, 6, 7, 8, 9, 10)
	fmt.Println(Number)
	fmt.Println(len(Number))
	fmt.Println(cap(Number))
	Number = append(Number, 11, 12, 13, 14, 15)
	fmt.Println(Number)
	fmt.Println(len(Number))
	fmt.Println(cap(Number))
}
