package main

import "fmt"

func main() {
	var revenue int
	var expenses int
	var taxRate float64

	print("Enter revenue: ")
	fmt.Scan(&revenue)

	print("Enter expenses: ")
	fmt.Scan(&expenses)

	print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := float64(revenue - expenses)
	profit := ebt * (1 - taxRate/100)
	fmt.Println("EBT:", ebt)
	fmt.Println("Profit after tax:", profit)

	ratio := profit / float64(revenue)
	// fmt.Println("Profit margin:", ratio*100, "%")
	// fmt.Printf("Profit margin: %.2f", ratio*100)

	// formattedFV := fmt.Sprintf("%.2f", ratio*100)
	// fomattedEBT := fmt.Sprintf("%.2f", ebt)
	// fomattedProfit := fmt.Sprintf("%.2f", profit)
	// fmt.Print(formattedFV, " ", fomattedEBT, " ", fomattedProfit)

	fmt.Printf(`Profit margin: %.2f
EBT: %.2f
Profit after tax: %.2f
	`, ratio*100, ebt, profit)

}
