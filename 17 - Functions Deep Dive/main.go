package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)

	// returning function from function
	rFunction := getReturnFunction()
	fmt.Println(rFunction(4))

	// anonymous function
	anonymousFunction := func(number int) int {
		return number * 2
	}
	fmt.Println(anonymousFunction(3))

	// Closures function
	double := closureFunction(2)
	triple := closureFunction(3)
	fmt.Println(double(5))
	fmt.Println(triple(5))

	// recursion function
	fmt.Println(recursionFactorial(5))

	// sum function
	fmt.Println(sumUp([]int{1, 2, 3}))

	// Variadic function
	fmt.Println(sumUpVariadic(1, 2, 3))
	// Variadic function for list
	numberList := []int{1, 2, 3}
	fmt.Println(sumUpVariadic(numberList...))
}

func doubleNumbers(n *[]int) []int {
	dNumbers := []int{}
	for _, val := range *n {
		// dNumbers = append(dNumbers, val*2)
		dNumbers = append(dNumbers, doubleSingleNumber(val))
	}
	return dNumbers
}

// returning function from function
func getReturnFunction() func(int) int {
	return doubleSingleNumber
}

func doubleSingleNumber(number int) int {
	return number * 2
}

// Closures functions
func closureFunction(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// recursion function
func recursionFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursionFactorial(n-1)
}

// sum function
func sumUp(list []int) int {
	sum := 0
	for _, val := range list {
		sum = sum + val
	}
	return sum
}

// Variadic Function
func sumUpVariadic(list ...int) int {
	sum := 0
	for _, val := range list {
		sum = sum + val
	}
	return sum
}
