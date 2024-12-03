package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank")
	// for i := 0; i < 3; i++ {
	// 	myBank()
	// }

	for {
		fmt.Println("What do you want to do?")
		fmt.Println(`
		1. Check balance
		2. Deposit money
		3. Withdraw money
		4. Exit`)

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		if choice == 1 {
			fmt.Println("The current Balance is: ", accountBalance)
		} else if choice == 2 {
			var depositInput float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositInput)
			if depositInput <= 0 {
				fmt.Println("Invalid input amount, must be greater than 0")
				continue
			}

			accountBalance += depositInput
			fmt.Println("Successful! The current Balance is: ", accountBalance)
		} else if choice == 3 {
			var withdrawInput float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawInput)
			if withdrawInput <= 0 {
				fmt.Println("Invalid input amount, must be greater than 0")
				continue
			}

			if withdrawInput > accountBalance {
				fmt.Println("Invalid amount, can't withdraw money")
				continue
			}

			accountBalance -= withdrawInput
			fmt.Println("Successful! The current balance is: ", accountBalance)
		} else if choice == 4 {
			fmt.Println("Thank You!")
			break
		} else {
			fmt.Println("Invalid input by user")
		}
	}
}

// switch choice {
// 	case 1:
// 		fmt.Println("Your balance is: ", accountBalance)
// 	case 2:
// 		var depositInput float64
// 		fmt.Print("Enter amount to deposit: ")
// 		fmt.Scan(&depositInput)
// 		if depositInput <= 0 {
// 			fmt.Println("Invalid input amount, must be greater than 0")
// 			continue
// 		}

// 		accountBalance += depositInput
// 		fmt.Println("Successful! The current Balance is: ", accountBalance)
// 	default:
// 		fmt.Println("GoodBye")
// 		return
// }
