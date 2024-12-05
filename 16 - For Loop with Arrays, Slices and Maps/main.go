package main

import "fmt"

func main() {
	fmt.Println("For loop with arrays, slices and maps")

	// for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Hello World")
	}

	// loop through array
	myArray := []int{1, 2, 3, 4, 5}
	for i := range myArray {
		fmt.Println(myArray[i], i)
	}

	// loop through maps
	myMap := map[string]int{
		"macOS":   10,
		"Linux":   9,
		"Windows": 5,
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
