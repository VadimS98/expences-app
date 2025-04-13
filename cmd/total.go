package cmd

import (
	"expenses-app/models"
	"fmt"
)

func ShowTotalExpenses(expenses []models.Expense) {
	total := 0.0
	for _, expense := range expenses {
		total += expense.Price * expense.Amount
	}

	fmt.Printf("Total Expenses: $%.2f\n", total)
}
