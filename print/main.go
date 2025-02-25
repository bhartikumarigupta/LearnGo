package main

import "fmt"

func main() {
	var name string = "Bharti"
	var rollno int = 76
	var marks float32 = 100.8
	fmt.Println("Name is", name, "Rollno is", rollno, "Marks is", marks)
	fmt.Printf("Name is %s Rollno is %d Marks is %.2f", name, rollno, marks)
	fmt.Println()
	fmt.Printf("Name is %T Rollno is %T Marks is %T", name, rollno, marks)
	fmt.Println()
	fmt.Println("Type of name is", name, "Type of rollno is", rollno, "Type of marks is", marks)

}
