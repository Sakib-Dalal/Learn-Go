package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata" // import package
)

func main() {
	fmt.Println(randomdata.SillyName())   // print silly name
	fmt.Println(randomdata.PhoneNumber()) // print dummy phone number

	fmt.Println(randomdata.Number(3))      // random number between 0 to 3
	fmt.Println(randomdata.Number(10, 12)) // random number between 10 to 12
}
