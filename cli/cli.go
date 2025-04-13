package cli

import (
	"bufio"
	"expenses-app/models"
	"fmt"
	"os"
	"strings"
)

func ShowMenu() {
	fmt.Println("Expenses App")
	fmt.Println("1. Add an Expense")
	fmt.Println("2. List All Expenses")
	fmt.Println("3. Show Total Expenses")
	fmt.Println("4. Exit")
}

func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// GetCategoryMenu displays a numbered list of categories and allows the user to select one
func GetCategoryMenu() string {
	fmt.Println("Available categories:")
	for i, category := range models.Categories {
		fmt.Printf("%d) %s\n", i+1, category)
	}

	var choice int
	for {
		fmt.Print("Enter the number corresponding to your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > len(models.Categories) {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		break
	}

	return models.Categories[choice-1]
}
