package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Learn File Handling in Go")
	// create file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	// content := "Hello Bharti let's start the journey"
	// _, errors := io.WriteString(file, content)
	// if errors != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return
	// }

	// file.WriteString(content)
	fmt.Println("File Created Successfully")
	file1, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file1.Close()
	// read file
	/*
		buffer := make([]byte, 1024)
		for {
			n, err := file1.Read(buffer)
			if err == io.EOF {
				// fmt.Println("End of file")
				break
			}
			if err != nil && err != io.EOF {
				fmt.Println("Error reading file:", err)
				return
			}
			fmt.Println(string(buffer[:n]))
		}
	*/
	// read the entire file into a byte slice
	content, err := ioutil.ReadAll(file1)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))

}
