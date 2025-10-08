package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Revenue : ")
	expenses := getUserInput("Expenses : ")
	taxRate := getUserInput("taxRate : ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Earning before tax :", ebt)
	fmt.Println("Profit :", profit)
	fmt.Println("Ratio : ", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
