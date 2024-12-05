## Functions Deep Dive

- Using Functions as Values & as type
```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)
}

func doubleNumbers(n *[]int) []int {
	dNumbers := []int{}
	for _, val := range *n {
		// dNumbers = append(dNumbers, val*2)
		dNumbers = append(dNumbers, doubleSingleNumber(val))
	}
	return dNumbers
}

func doubleSingleNumber(number int) int {
	return number * 2
}
```

- Anonymous Functions
```go
// anonymous function
anonymousFunction := func(number int) int {
	return number * 2
}
fmt.Println(anonymousFunction(3))
```

- Recursion
```go
// recursion function call in main
fmt.Println(recursionFactorial(5))

//  recursion function
func recursionFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursionFactorial(n-1)
}
```

- Variadic Function
```go
// Variadic function call in main
fmt.Println(sumUpVariadic(1, 2, 3))
// Variadic function for list call in main
numberList := []int{1, 2, 3}
fmt.Println(sumUpVariadic(numberList...))

// Variadic Function
func sumUpVariadic(list ...int) int {
	sum := 0
	for _, val := range list {
		sum = sum + val
	}
	return sum
}
```