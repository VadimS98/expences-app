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
	// Get user input for description
	description := cli.GetUserInput("Enter description: ")

	// Get user input for category with validation using models.Categories
	category := getValidCategory()

	// Get user input for price with validation
	price := getValidFloatInput("Enter price: ")

	// Get user input for amount with validation
	amount := getValidFloatInput("Enter amount: ")

	// Create and append the expense
	expense := models.Expense{
		Price:       price,
		Amount:      amount,
		Category:    category,
		Description: description,
	}
	expenses = append(expenses, expense)
	fmt.Println("Expense added!")
}

// Helper function to validate category input using models.Categories
func getValidCategory() string {
	for {
		fmt.Println("Available categories:")
		for _, cat := range models.Categories {
			fmt.Println("-", cat)
		}

		category := cli.GetUserInput("Choose a category: ")
		for _, cat := range models.Categories {
			if category == cat {
				return category
			}
		}

		fmt.Println("Invalid category. Please choose from the list.")
	}
}

// Helper function to validate float input (e.g., price, amount)
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
