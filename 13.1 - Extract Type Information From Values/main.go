package main

import "fmt"

func main() {
	printValueType(2.4)
}

func printValueType(value interface{}) {
	// fmt.Println(value.(int))

	intVal, ok := value.(int)
	if ok {
		fmt.Printf("value is: %v and it is int?: %v\n", intVal, ok)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Printf("value is: %v and it is float?: %v\n", floatVal, ok)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Printf("value is: %v and it is float?: %v\n", stringVal, ok)
		return
	}

}
