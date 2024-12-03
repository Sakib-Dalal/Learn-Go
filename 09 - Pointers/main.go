package main

import "fmt"

func main() {

	// Pointer address example
	age := 32
	agePointer := &age
	// or
	// var agePointer *int
	// agePointer = &age
	fmt.Printf("The value is %d and it's address is %v, age value of pointer is %v\n", age, agePointer, *agePointer)

	// null pointer
	var nullPointer *int
	fmt.Println(nullPointer)

	userAge := 32
	// function without pointers
	fmt.Println("Age: ", age)
	adultYears := getAdultYears(userAge)
	fmt.Println(adultYears)

	// using pointer and passing to function
	userAgePointer := &userAge
	adultYearsPointer := getAdultYearsUsingPointer(userAgePointer)
	fmt.Println(adultYearsPointer)

}

// no pointer function
func getAdultYears(age int) int {
	return age - 18
}

// pointer function
func getAdultYearsUsingPointer(age *int) int {
	return *age - 18
}
