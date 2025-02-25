/*varriable declaration in go lang and printing a statement in go lang*/

// package main

// import "fmt"

// func main() {
// 	var username string = "John Doe"
// 	fmt.Println("Username is ", username)
// 	fmt.Printf("Username is %s\n", username)
// 	fmt.Printf("Username type  is %v\n", username)
// 	fmt.Printf("Username type  is %T\n", username)
// }

/* take input using  go from user*/
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Print("Enter your name: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	//comma ok /error ok
// 	username, _ := reader.ReadString('\n')
// 	fmt.Println("Username is ", username)
// 	fmt.Printf("Username is %T\n", username)
// }

/* Conversion of data types in go lang*/
package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

func main() {

	fmt.Printf("Welcome To My Pizza Hut\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please give rating to our pizza:\n ")
	// rating, _ := reader.ReadString('\n')
	fmt.Printf("Initially, Rating type is %T\n", reader)
	// //coversion part
	// newrating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	// if err != nil {
	// 	fmt.Println("Error while converting rating to float64", err)
	// } else {
	// 	// Fmt.Println("Rating after conversion ",rating)
	// 	fmt.Printf("Rating type is %T\n", newrating)
	// }
	// There are total 25 keywords present in the Go language as follows:

	// break
	// case
	// chan
	// const
	// continue
	// default
	// defer
	// else
	// fallthrough
	// for
	// func
	// go
	// goto
	// if
	// import
	// interface
	// map
	// package
	// range
	// return
	// select
	// struct
	// switch
	// type
	// var
	/*time in go lang */
	// presenttime := time.Now()
	// fmt.Println("Present time is ", presenttime)
	// fmt.Println("Present time is ", presenttime.Format("02-01-2006 "))
	// fmt.Println("Present time is ", presenttime.Format("01-02-2006 03:04:05 PM"))
	// fmt.Println("Present time is ", presenttime.Format("01-02-2006 03:04:05 PM -0700"))
	// fmt.Println("Present time is ", presenttime.Format("01-02-2006 03:04:05 PM -0700 MST"))
	// fmt.Println("Present time is ", presenttime.Format("01-02-2006 03:04:05 PM -0700 MST"))
	// fmt.Println("month is ", presenttime.Month())
	// fmt.Println("day is ", presenttime.Day())
	// fmt.Println("year is ", presenttime.Year())
	// fmt.Println("hour is ", presenttime.Hour())
	// fmt.Println("minute is ", presenttime.Minute())
	// fmt.Println("second is ", presenttime.Second())
	// fmt.Println("week day is ", presenttime.Weekday())
	// fmt.Println("week day is ", presenttime.Weekday().String())
	// fmt.Println("week day is ", presenttime.Weekday().String()[:3])
	// fmt.Println("week day is ", presenttime.Weekday().String()[:3])
	/*build in windows*/
	// fmt.Println("GOOS: ", runtime.GOOS)
	// GOOSS := "windows" go build  (run this command in terminal)

}

/* Memory management in go lang*/
// Memory allocation and deallocation in go lang happens automatically
/* new()
allocate memory but no Init
you will get a memory address
zeroed storage
*/
/*
make()
allocate memory and Init
you will get a memory address
non-zeroed storage
*/
// Garbage collection in go lang is automatic
// Out of Scope or nil

// POINTERS IN GO LANG
/*
package main

import "fmt"

func main() {
	var num1 int = 10
	var ptr = &num1
	fmt.Println("Value of num1 is ", num1)
	//address of num1 or pointer of num1
	fmt.Println("Address of num1 is ", ptr)
	//value of num1
	fmt.Println("Value of num1 is ", *ptr)
}
*/
//Arrays in go lang
/*
package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	arr[5] = 60
	arr[6] = 70
	arr[7] = 80
	arr[8] = 90
	arr[9] = 100
	fmt.Println("Array is ", arr)
	fmt.Println("Array is ", arr[0])
	fmt.Println("Array is ", arr[1])
	var str [5]string
	str[0] = "John"
	str[1] = "Doe"
	str[2] = "Jane"
	str[3] = "Doe"
	str[4] = "Smith"
	fmt.Println("String Array is ", str)
	var arr1 = [5]string{"John", "Doe", "Jane"}
	fmt.Println("String Array is ", len(arr1))
}
*/
//Slices in go lang
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// var slicearr = []string{"John", "Doe", "Jane"}
// fmt.Println("String Array is ", slicearr)
// slicearr = append(slicearr, "Smith", "Alex")
// fmt.Println("String Array is ", slicearr)
// fmt.Println("String Array is ", len(slicearr))
// fmt.Println("String Array is ", cap(slicearr))
// //slicing
// fmt.Println("String Array is ", slicearr[1:3])
// fmt.Println("String Array is ", slicearr[:3])
// fmt.Println("String Array is ", slicearr[1:])
// fruitlist := make([]int, 5)
// fruitlist[0] = 5
// fruitlist[1] = 2
// fruitlist[2] = 4
// fruitlist[3] = 1
// fruitlist[4] = 3
// fmt.Println("Fruit List is ", fruitlist)
// fmt.Println("Fruit List is ", len(fruitlist))
// // fruitlist[5] = 6 //this line show error because we have only 5 elements in fruitlist but we can reallocate the memory using append function
// // fmt.Println("Fruit List is ", fruitlist)
// fruitlist = append(fruitlist, 6, 7, 8, 9, 10)
// fmt.Println("Fruit List is ", fruitlist)
// fmt.Println("Fruit List is ", len(fruitlist))
// //to check slices are sorted or not
// fmt.Println("Is Fruit List sorted ", sort.IntsAreSorted(fruitlist))
// //sort the slice
// sort.Ints(fruitlist)
// fmt.Println("Fruit List after sorting is ", fruitlist)
// fmt.Println("Is Fruit List sorted ", sort.IntsAreSorted(fruitlist))
// //how to remove element from slice based on index
// var index int = 2
// fruitlist = append(fruitlist[:index], fruitlist[index+1:]...)
// fmt.Println("Fruit List after removing element is ", fruitlist)

// }
// Maps in go lang
/*
package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["en"] = "English"
	languages["fr"] = "French"
	languages["es"] = "Spanish"
	languages["de"] = "German"
	languages["it"] = "Italian"
	fmt.Println("Languages are ", languages)
	fmt.Println("Languages are ", languages["en"])
	//delete element from map
	delete(languages, "en")
	fmt.Println("Languages are ", languages)
	//check if key is present in map or not
	lang, ok := languages["en"]
	if ok {
		fmt.Println("Language is ", lang)
	} else {
		fmt.Println("Language not found")
	}
	//iterate over map
	for key, value := range languages {
		fmt.Println("Key is ", key, " Value is ", value)
	}
	//or
	for _, value := range languages {
		fmt.Println("Value is ", value)
	}

}
*/
// struct in go lang
/*package main

import "fmt"

type User struct {
	Username string
	Age      int
	Email    string
	Address  string
	Grade    string
}

func main() {
	User1 := User{
		"Bharti Gupta", 22, "shadhviraj97099@gmail.com", "Patna", "+A"}
	fmt.Println("User is ", User1)
	fmt.Println("User  name is ", User1.Username)
	fmt.Println("User  age is ", User1.Age)
	fmt.Println("User  email is ", User1.Email)
	fmt.Println("User  address is ", User1.Address)
	fmt.Println("User  grade is ", User1.Grade)
	fmt.Printf("User  detail is %+v\n", User1)
}
*/
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Println("please enter your average percentage")
// 	reader := bufio.NewReader(os.Stdin)
// 	percentage, _ := reader.ReadString('\n')
// 	cnvrtpercentage, err := strconv.ParseFloat(strings.TrimSpace(percentage), 64)
// 	if err != nil {
// 		fmt.Println("Error while converting percentage to float64", err)
// 	} else {
// 		fmt.Println("Percentage is ", cnvrtpercentage)
// 		if cnvrtpercentage >= 70 {
// 			fmt.Println("You are passed with first division")
// 		} else if cnvrtpercentage >= 30 && cnvrtpercentage < 70 {
// 			fmt.Println("You are passed with second division")
// 		} else {
// 			fmt.Println("You are failed")

// 		}
// 	}
// 	if num := 10; num%2 == 0 {
// 		fmt.Println("Number is even")
// 	} else {
// 		fmt.Println("Number is odd")
// 	}
// 	// fmt.Scanln(&percentage)

// }
/* output :-please enter your average percentage
40
Percentage is  40
You are passed with second division
Number is even  */
//Switch case in go lang
/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Intn(5) + 1
	fmt.Println("Random number is ", num)
	switch num {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
	case 4:
		fmt.Println("Number is 4")
	case 5:
		fmt.Println("Number is 5")
	default:
		fmt.Println("Number is not in range")
	}
}
*/
// for loop in go lang
// package main

// import "fmt"

//	func main() {
//		arr := []int{10, 20, 30, 40, 50}
//		// for i := 0; i < len(arr); i++ {
//		// 	fmt.Println("Element at ", i, " is ", arr[i])
//		// }
//		// for i := range arr {
//		// 	fmt.Println("Element at ", i, " is ", arr[i])
//		// }
//		for i, value := range arr {
//			// if arr[i] == 40 {
//			// 	fmt.Println("Element found at ", i)
//			// 	break
//			// }
//			// if arr[i] == 20 {
//			// 	fmt.Println("Element found at ", i)
//			// 	continue
//			// }
//			if arr[i] == 30 {
//				goto lco
//			}
//			fmt.Println("Element at ", i, " is ", value)
//		}
//
// lco:
//
//		fmt.Println("Lets Code Online")
//	}
//
// functions in go lang
/*
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the greeting world")
	greet()
	//this is not valid syntex of writing function inside function in go lang
	// func greet1(){
	// 	fmt.Println("Hello, Good Morning")
	// }
	result := addition(10, 20)
	fmt.Println("Addition is ", result)
	result1 := concat("Hello", "World")
	fmt.Println("Concatenation is ", result1)
	result2 := proadder(10, 20, 30, 40, 50)
	fmt.Println("Addition is ", result2)
	result3, result4 := calculation(10, 20, 30, 40, 50)
	fmt.Println("Addition is ", result3)
	fmt.Println("Product is ", result4)

}
func greet() {
	fmt.Println("Hello, Good Morning")
}
func addition(num1 int, num2 int) int {
	return num1 + num2
}
func concat(str1 string, str2 string) string {
	return str1 + str2
}
func proadder(num1 ...int) int {
	sum := 0
	for _, value := range num1 {
		sum += value
	}
	return sum

}
func calculation(num1 ...int) (int, int) {
	sum := 0
	product := 1
	for _, value := range num1 {
		sum += value
		product *= value
	}
	return sum, product

}
*/
//Methods in go lang
// package main

// import "fmt"

// func main() {
// 	user1 := User{
// 		Username: "John Doe",
// 		Age:      30,
// 		Email:    "shadhviraj97099@gmail.com",
// 		Address:  "Patna",
// 	}
// fmt.Println("User is ", user1)
// fmt.Println("User  name is ", user1.Username)
// fmt.Println("User  age is ", user1.Age)
// fmt.Println("User  email is ", user1.Email)
// fmt.Println("User  address is ", user1.Address)
// 	fmt.Println(user1.greet())
// 	user1.NewEmail()
// 	fmt.Println("User  email is ", user1.Email)
// }

// type User struct {
// 	Username string
// 	Age      int
// 	Email    string
// 	Address  string
// }

// func (u User) greet() string {
// 	return "Hello, Good Morning " + u.Username
// }
// func (u User) NewEmail() {
// 	u.Email = "bharti@123gmai.com"
// 	println("New Email is ", u.Email)
// }
//Defer in go lang
/*  A defer statement defers the execution of a function until the surrounding function returns.
A Defer  statement invokes a function whose execution is deferred until the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.*/
// package main

// import "fmt"

// func main() {
// defer fmt.Println("Hello")
// fmt.Println("World")// output:-World Hello
// defer fmt.Println("one")
// defer fmt.Println("two")
// defer fmt.Println("three")
// fmt.Println("four") //output:-four three two one
// defer fmt.Println("one")
// fmt.Println("two")
// defer fmt.Println("three")
// fmt.Println("four") //output:-two four three one
// defer fmt.Println("one")
// fmt.Println("two")
// fmt.Println("three") //output:-two three one
// }
// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello")
// 	defer fmt.Println("World")
// 	fmt.Println("Lets Code Online")
// 	defer fmt.Println("Welcome")
// 	defer fmt.Println("To")
// 	defer fmt.Println("The")
// 	// defer fmt.Println("World")
// 	MyDefer()

// }
//
//	func MyDefer() {
//		for i := 0; i < 5; i++ {
//			defer fmt.Println(i)
//		}
//	}  //output:-Hello
//
// Lets Code Online
// 4
// 3
// 2
// 1
// 0
// The
// To
// Welcome
// World
// working with files in go lang
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	fmt.Printf("Welcome to the file handling in go lang\n")
// 	content := "Hello, Welcome to the file handling in go lang"
// 	file, err := os.Create("file.txt")
// 	// if err != nil {
// 	// 	panic(err)
// 	// 	// fmt.Println("Error while creating file", err)l
// 	// }
// 	checkerror(err)
// 	length, err := file.WriteString(content)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	checkerror(err)
// 	fmt.Println("Length of content is ", length)
// 	file.Close()
// 	readfile("file.txt")
// }
// func readfile(filename string) {
// 	databyte, err := ioutil.ReadFile(filename)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	checkerror(err)
// 	fmt.Println("Text Data in file is ", string(databyte))
// }
// func checkerror(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// Handaling web request in go lang
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// const url = "https://jsonplaceholder.typicode.com/posts"

//	func main() {
//		fmt.Println("Welcome to the web request in go lang")
//		response, err := http.Get(url)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("Response is ", response.StatusCode)
//		fmt.Printf("Response is %T\n", response)
//		defer response.Body.Close() //caller responsibility to close the connection
//		// fmt.Println("Response is ", response.Status)
//		// fmt.Println("Response is ", response.Proto)
//		// fmt.Println("Response is ", response.ProtoMajor)
//		// fmt.Println("Response is ", response.ProtoMinor)
//		// fmt.Println("Response is ", response.Header)
//		// fmt.Println("Response is ", response.ContentLength)
//		// fmt.Println("Response is ", response.TransferEncoding)
//		// fmt.Println("Response is ", response.Close)
//		databytes, err := ioutil.ReadAll(response.Body)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("Data is ", string(databytes))
//	}
// package main

// import (
// 	"fmt"
// 	"net/url"
// )

// const myurl string = "https://leetcode.com/u/burnwal13ankita07/"

// func main() {
// 	fmt.Printf("Welcome to the hadling url request in go lang\n")
// 	response, err := url.Parse(myurl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Response is ", response)
// 	fmt.Println("Response is ", response.Scheme)
// 	fmt.Println("Response is ", response.Host)
// 	fmt.Println("Response is ", response.Path)
// 	fmt.Println("Response is ", response.Port())
// 	qryprm := response.Query()
// 	fmt.Printf("Query Parameters are %+v\n", qryprm)
// 	fmt.Println("Query Parameters are ", qryprm["u"])
// 	fmt.Println("Query Parameters are ", qryprm["u"][0])
// 	fmt.Printf("Query Parameters are %T\n", qryprm)
// }
