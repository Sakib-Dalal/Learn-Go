package main

import (
	"fmt"
	"time"
	// "example.com/structs_app/mystructure"
)

// create a structure
type myStruct struct {
	userFirstName string
	userLastName  string
	userBirthdate string
	createdAt     time.Time
}

// function within struct
func (user myStruct) outputUserDetails2() {
	fmt.Println(user)
	fmt.Println(user.userFirstName, user.userLastName, user.userBirthdate)
}

func (user *myStruct) clearUserDetails() {
	user.userFirstName = " "
	user.userLastName = " "
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

	// function within structure
	appUser.outputUserDetails2()

	// access fields of the struct
	fmt.Println(appUser.userFirstName)
	fmt.Println(appUser.userLastName)
	fmt.Println(appUser.userBirthdate)

	appUser.clearUserDetails()
	appUser.outputUserDetails2()

	// structure from other file module
	// newUser := mystructure.UserStructure{
	// 	FirstName: "jay",
	// 	LastName:  "roy",
	// 	Age:       "45",
	// }
	// fmt.Println(newUser)

}

// passing struct to function
func outputUserDetails(user myStruct) {
	fmt.Println(user)
	fmt.Println(user.userFirstName, user.userLastName, user.userBirthdate)
}

func getUserData(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}
