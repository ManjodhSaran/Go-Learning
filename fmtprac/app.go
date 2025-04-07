package main

import (
	"fmt"
)

func main() {
	var investment float64 = 10000
	var interestRate float64
	var years int

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investment)
	fmt.Println("investment:", investment)

	fmt.Print("Enter interest rate: ")
	fmt.Scan(&interestRate)
	fmt.Println("interest rate:", interestRate)
	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)
	fmt.Println("years:", years)

	var principal = investment
	var interest = 0.0
	for i := 0; i < years; i++ {
		interest = interest + principal*interestRate
		principal = principal * (1 + interestRate)
	}
	fmt.Println("principal:", principal)
	fmt.Println("interest:", interest)
}
