package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Investment Calculator")

	const inflationRate float64 = 2.5
	var nullVariable float64
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5

	// over write investmentAmount
	investmentAmount = 2000

	// get input from the user - add pointer to the variable to take as input
	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount*math.Pow(1+expectedReturnRate/100, years) + nullVariable
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
