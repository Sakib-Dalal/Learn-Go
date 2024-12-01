## Functions
- main function ``` func main() {} ```
- other function's
```
func main() {
	fmt.Println("This is main function")

	// call simple function
	simpleFunction()

	// call function with parameter
	age := 34
	name := "roy"
	paramFunction(age, name)
}

// simple function
func simpleFunction() {
	fmt.Println("This is simple function")
}

// parameter function
func paramFunction(age int, name string) {
	fmt.Printf("The input age is %d and name is %s \n", age, name)
}
```

- return function
``` 
func main() {
	fmt.Println("This is main function")

	// return value function
	answer := additionFunction(45, 5)
	fmt.Printf("The answer is %d.\n", answer)
}

// return value function
func additionFunction(a, b int) int {
	return a + b
}
```

- return multiple value's
```
func main() {
	fmt.Println("This is main function")

	// return multiple values function
	multiply, divide := divideAndMultiply(45, 5)
	fmt.Printf("The multiplication is %.2f and division is %.2f.\n", multiply, divide)
}

// return multiple values function
func divideAndMultiply(a float64, b float64) (float64, float64) {
	return (a * b), (a / b)
}
```