package cli

import (
	"bufio"
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
