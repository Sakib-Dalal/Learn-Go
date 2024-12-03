package main

import (
	"fmt"

	"example.com/module/datetime" // import datetime package
)

func welcomeFunction() {
	fmt.Println("Hello World! Good Morning")
	// say data time module from datetime/data.go package
	datetime.SayDateTime()
}
