package cmd

import (
	"expenses-app/models"
	"fmt"
)

func ListExpenses(expenses []models.Expense) {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}

	fmt.Println("Expenses:")
	for i, expense := range expenses {
		fmt.Printf("%d. %s -> %s -> %.2f -> $%.2f\n", i+1, expense.Description, expense.Category, expense.Amount, expense.Price)
	}
}
