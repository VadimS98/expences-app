package cmd

import (
	"expenses-app/cli"
	"expenses-app/internal"
	"expenses-app/models"
	"fmt"
)

func AddExpense(expenses *[]models.Expense) {
	description := cli.GetUserInput("Enter description: ")
	category := cli.GetCategoryMenu()
	price := internal.GetValidFloatInput("Enter price: ")
	amount := internal.GetValidFloatInput("Enter amount: ")

	expense := models.Expense{
		Description: description,
		Category:    category,
		Price:       price,
		Amount:      amount,
	}

	*expenses = append(*expenses, expense)
	fmt.Println("Expense added!")
}
