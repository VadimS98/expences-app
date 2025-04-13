package cmd

import (
	"expenses-app/cli"
	"expenses-app/models"
	"fmt"
	"strconv"
)

func AddExpense(expenses *[]models.Expense) {
	description := cli.GetUserInput("Enter description: ")
	category := cli.GetCategoryMenu()
	price := getValidFloatInput("Enter price: ")
	amount := getValidFloatInput("Enter amount: ")

	expense := models.Expense{
		Description: description,
		Category:    category,
		Price:       price,
		Amount:      amount,
	}

	*expenses = append(*expenses, expense)
	fmt.Println("Expense added!")
}

func getValidFloatInput(prompt string) float64 {
	for {
		input := cli.GetUserInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil || value < 0 {
			fmt.Println("Invalid input. Please enter a positive number.")
			continue
		}
		return value
	}
}
