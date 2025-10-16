package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, err := getUserInput("Revenue : ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("________")
		return
	}

	expenses, err := getUserInput("Expenses : ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("________")
		return
	}

	taxRate, err := getUserInput("taxRate : ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("________")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax : %.1f \n", ebt)
	fmt.Printf("Profit : %.1f \n", profit)
	fmt.Printf("Ratio : %.3f \n", ratio)
	storeResultIntoFile(ebt, profit, ratio)
}

func storeResultIntoFile(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT :%.1f\nProfit :%.1f\nRatio :%.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(result), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return userInput, errors.New("value must be bigger than 0")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
