package helpers

import (
	"errors"
	"fmt"
	"os"
)

func StoreResultIntoFile(ebt, profit, ratio float64) {
	result := fmt.Sprintf("EBT :%.1f\nProfit :%.1f\nRatio :%.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(result), 0644)
}

func GetUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return userInput, errors.New("value must be bigger than 0")
	}

	return userInput, nil
}

func CalculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
