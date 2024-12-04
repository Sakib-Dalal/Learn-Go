package main

import "fmt"

func main() {
	myArray := [4]float64{1.2, 2.2, 3.1, 5.6} // create float array // store 4 values
	fmt.Println(myArray)

	prices := [10]int{1, 2, 3}
	fmt.Println(prices)

	var productName [4]string = [4]string{"di", "go"}
	fmt.Println(productName)

	// access value by index in array
	fmt.Println(prices[0])
	fmt.Println(prices[5])
	fmt.Println(prices[1])

	// set value of index in array
	prices[5] = 50
	fmt.Println(prices)

	// slicing in arrays
	fmt.Println(prices[2:6])
	fmt.Println(prices[:4])
	fmt.Println(prices[4:])

	// get array length and capacity
	fmt.Println(len(prices), cap(prices))

	// dynamic arrays // auto scalable array size
	dynamicArray := []int{1, 23, 3}
	fmt.Println(dynamicArray)
	// add new value in dynamic array
	updatedDynamicArray := append(dynamicArray, 4)
	fmt.Println(updatedDynamicArray, dynamicArray)
	// reassign array
	dynamicArray = updatedDynamicArray
	fmt.Println(dynamicArray)
}