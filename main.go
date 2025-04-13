package main

import (
	"expenses-app/cmd"
	"expenses-app/models"
	"expenses-app/storage"
	"fmt"
)

const dataFile = "expenses.json"

func main() {
	var expenses []models.Expense

	expenses, err := storage.LoadExpensesFromFile(dataFile)
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	defer func() {
		if err := storage.SaveExpensesToFile(expenses, dataFile); err != nil {
			fmt.Println("Error saving data:", err)
		}
	}()

	cmd.Run(&expenses, dataFile)
}
