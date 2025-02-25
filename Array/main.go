package main

import "fmt"

func main() {
	fmt.Println("We are learning Arrays in Go")
	var name [5]string
	name[0] = "Bharti"
	name[1] = "Raj"
	name[2] = "Rahul"
	name[3] = "Ravi"
	name[4] = "Rohit"
	fmt.Println(name)
	fmt.Println(name[0])
	// fmt.Println(name[5]) // index out of range
	fmt.Println(len(name)) // length of array
	fmt.Println(name[0:2]) // slicing of array
	var city = [5]string{"Delhi", "Mumbai", "Kolkata", "Chennai", "Bangalore"}
	fmt.Println(city)
	fmt.Println(city[0:2]) // will print Delhi, Mumbai
	fmt.Println(city[0:3]) // will print Delhi, Mumbai, Kolkata
	fmt.Println(city[2:4]) // will print Kolkata, Chennai
	fmt.Println(city[2:])  // will print Kolkata, Chennai, Bangalore
	fmt.Println(city[:3])  // will print Delhi, Mumbai, Kolkata
	fmt.Printf("%q\n", city)

}
