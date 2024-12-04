package main

import (
	"fmt"
)

func main() {
	ansValue := addTwoValues(2, 2)
	fmt.Println(ansValue)

	ansValue = addTwoValues(2.1, 1.2)
	fmt.Println(ansValue)

	ansValue = addTwoValues("s", "d")
	fmt.Println(ansValue)

	ansValue = addTwoValues(2, "d")
	fmt.Println(ansValue)
}

func addTwoValues(a, b interface{}) interface{} {
	// for integer
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	// for float
	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	// for string
	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}

	return nil
}
