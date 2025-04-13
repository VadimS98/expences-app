package cmd

import (
	"expenses-app/cli"
	"expenses-app/models"
	"fmt"
)

func Run(expenses *[]models.Expense, dataFile string) {
	for {
		cli.ShowMenu()
		choice := cli.GetUserInput("Choose an option: ")

		switch choice {
		case "1":
			AddExpense(expenses)
		case "2":
			ListExpenses(*expenses)
		case "3":
			ShowTotalExpenses(*expenses)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
