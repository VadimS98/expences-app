package main

import (
	"expenses-app/cmd"
	"expenses-app/internal"
	"expenses-app/models"
	"fmt"
)

const dataFile = "expenses.json"

func main() {
	var expenses []models.Expense

	expenses, err := internal.LoadExpenses(dataFile)
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	defer func() {
		if err := internal.SaveExpenses(expenses, dataFile); err != nil {
			fmt.Println("Error saving data:", err)
		}
	}()

	cmd.Run(&expenses, dataFile)
}
