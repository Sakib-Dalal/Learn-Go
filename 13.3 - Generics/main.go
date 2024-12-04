package main

import "fmt"

func main() {
	result := addValues(1, 2)
	fmt.Println(result)
}

func addValues[T int | float64 | string](a, b T) T {
	return a + b
}
