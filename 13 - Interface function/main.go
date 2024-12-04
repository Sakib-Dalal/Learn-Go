package main

import (
	"fmt"
)

func main() {
	// print any type of value
	printAnyValue(1)
	printAnyValue(1.5)
	printAnyValue("s")
	printAnyValue(true)

	printType(1)

	// another approach using switch
	printAnyValueUsingSwitch(1)
	printAnyValueUsingSwitch(1.2)
	printAnyValueUsingSwitch("s")
}

func printAnyValue(value interface{}) {
	fmt.Println(value)
}

func printType(value interface{}) {
	fmt.Printf("%T\n", value)
}

func printAnyValueUsingSwitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}
}
