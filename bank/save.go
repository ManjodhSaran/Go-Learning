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
