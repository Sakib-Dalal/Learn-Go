## Structs
- Used for grouping data and functions, methods into collections

- example 
- https://golangbot.com/structs/

- struct example
```go
type myStruct struct {
    // variables here with type
}{
    // define variable data here
}
```

- example code:
```go
package main

import (
	"fmt"
	"time"
)

// create a structure
type myStruct struct {
	userFirstName string
	userLastName  string
	userBirthdate string
	createdAt     time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	fmt.Println(firstName, lastName, birthdate)

	// create a variable for structure
	// var appUser myStruct
	appUser := myStruct{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthdate: birthdate,
		createdAt:     time.Now(),
	}
	// empty struct: appUser = myStruct {}
	// fmt.Println(appUser)
	outputUserDetails(appUser)

	// access fields of the struct
	fmt.Println(appUser.userFirstName)
	fmt.Println(appUser.userLastName)
	fmt.Println(appUser.userBirthdate)
}

// passing struct to function
func outputUserDetails(user myStruct) {
	fmt.Println(user)
}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)
	return value
}
```