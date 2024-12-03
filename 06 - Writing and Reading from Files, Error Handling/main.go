package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var data string = "Hello World"

	// write data in text file
	os.WriteFile("go_text.txt", []byte(data), 0644)

	// read data from text file
	new_data, _ := os.ReadFile("go_text.txt") // returns two values
	fmt.Println(new_data)

	read_data := string(new_data) // convert to string
	fmt.Println(read_data)

	float_data, _ := strconv.ParseFloat(read_data, 64) // convert back to float
	fmt.Println(float_data)

	// error handling
	data1, err := os.ReadFile("go.txt")
	if err != nil {
		// approach no 1
		fmt.Println("Error file reading") // if using return use (error)
		fmt.Println(err)
		fmt.Println("------------------------")
		// approach no 2
		// panic("File not found")
		panic(err)
	} else {
		fmt.Println(string(data1))
	}
}
