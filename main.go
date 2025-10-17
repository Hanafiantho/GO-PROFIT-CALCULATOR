package main

import (
	"fmt"

	"profit-calculator/helpers"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, err := helpers.GetUserInput("Revenue : ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("________")
		return
	}

	expenses, err := helpers.GetUserInput("Expenses : ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("________")
		return
	}

	taxRate, err := helpers.GetUserInput("taxRate : ")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("________")
		return
	}

	ebt, profit, ratio := helpers.CalculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax : %.1f \n", ebt)
	fmt.Printf("Profit : %.1f \n", profit)
	fmt.Printf("Ratio : %.3f \n", ratio)
	helpers.StoreResultIntoFile(ebt, profit, ratio)
}
