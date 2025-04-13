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

// GetCategoryInput prompts the user to choose a category from the available list
func GetCategoryInput() string {
	fmt.Println("Available categories:")
	for _, category := range models.Categories {
		fmt.Println("-", category)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose a category: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Validate input against available categories
	for _, category := range models.Categories {
		if input == category {
			return input
		}
	}

	fmt.Println("Invalid category. Please choose from the list.")
	return GetCategoryInput() // Retry on invalid input
}
