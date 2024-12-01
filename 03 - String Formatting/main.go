package main

import "fmt"

func main() {
	// integer
	num := 5
	fmt.Println("The number is ", num)
	fmt.Printf("This is number %v which is '%T' data type\n", num, num)
	fmt.Printf("This is number %v and square is %v \n", num, num*num)

	// float
	floatNum := 10.30393203243242345
	fmt.Printf("The float number is: %.2f \n", floatNum)
	fmt.Printf("The float number is '%T' data type \n", floatNum)

	// combined string sprint formatting
	stringVal := "Hello"
	combinedString := fmt.Sprintf("The number is %d, and float is %f and string is '%s'.\n", num, floatNum, stringVal)
	fmt.Print(combinedString)

	// multiline string
	fmt.Printf(`This is multiline string example: 
	int is - %d,
	float is - %f
	string is - %s 
	`, num, floatNum, stringVal)
}
