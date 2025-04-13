package main

import (
	"expenses-app/cli"
	"expenses-app/models"
	"expenses-app/storage"
	"fmt"
	"strconv"
)

var expenses []models.Expense

const dataFile = "expenses.json"

func main() {
	var err error
	expenses, err = storage.LoadExpensesFromFile(dataFile)
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	defer func() {
		if err := storage.SaveExpensesToFile(expenses, dataFile); err != nil {
			fmt.Println("Error saving data:", err)
		}
	}()

	for {
		cli.ShowMenu()
		choice := cli.GetUserInput("Choose an option: ")

		switch choice {
		case "1":
			addExpense()
		case "2":
			listExpenses()
		case "3":
			showTotalExpenses()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func addExpense() {
	description := cli.GetUserInput("Enter description: ")
	category := cli.GetUserInput("Enter category: ")
	priceStr := cli.GetUserInput("Enter price: ")
	price, _ := strconv.ParseFloat(priceStr, 64)
	amountStr := cli.GetUserInput("Enter amount: ")
	amount, _ := strconv.ParseFloat(amountStr, 64)

	expense := models.Expense{Price: price, Amount: amount, Category: category, Description: description}
	expenses = append(expenses, expense)
	fmt.Println("Expense added!")
}

func listExpenses() {
	if len(expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}

	fmt.Println("Expenses:")
	for i, expense := range expenses {
		fmt.Printf("%d. %s -> %s -> %.2f -> $%.2f\n", i+1, expense.Description, expense.Category, expense.Amount, expense.Price)
	}
}

func showTotalExpenses() {
	total := 0.0
	for _, expense := range expenses {
		total += expense.Price * expense.Amount
	}

	fmt.Printf("Total Expenses: $%.2f\n", total)
}
