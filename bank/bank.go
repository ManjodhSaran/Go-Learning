package main

import (
	"fmt"
	"os"
	"strconv"
)

const file = "bal.txt"

func getBalanceFromFile(defaultVal float64) float64 {

	if defaultVal == 0 {
		defaultVal = 1000
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return defaultVal
	}
	bal, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return defaultVal
	}
	return bal
}

func writeToFile(amt float64) {
	txt := fmt.Sprint(amt)
	os.WriteFile(file, []byte(txt), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()
	fmt.Println("Welcome to the Bank!")

	for {
		fmt.Println("----------------------")
		fmt.Println("What would you like to do?")
		fmt.Print("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit\n")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		fmt.Printf("%d Choosen\n", choice)

		switch choice {
		case 1:
			fmt.Println("Balance:", accountBalance)
		case 2:
			fmt.Print("Enter Amount to Deposit: ")
			var val float64
			fmt.Scan(&val)

			if val <= 0 {
				fmt.Println("Amount must be more than 0")
				continue
			}

			accountBalance += val
			writeToFile(accountBalance)
			fmt.Println("New Balance: ", accountBalance)
		case 3:
			fmt.Print("Enter Amount to withdraw: ")
			var val float64
			fmt.Scan(&val)

			if val <= 0 {
				fmt.Println("Amount must be more than 0")
				continue
			}

			if val > accountBalance {
				fmt.Println("InSufficient  funds")
				continue
			}

			accountBalance -= val
			writeToFile(accountBalance)
			fmt.Println("New Balance: ", accountBalance)
		case 4:
			fmt.Println("See you Again!")
			fmt.Println("Good-Bye!")
			return
		default:
			fmt.Println("Invalid Entry, Try Again")
		}
	}

}
