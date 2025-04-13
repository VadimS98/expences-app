package internal

import (
	"expenses-app/cli"
	"fmt"
	"strconv"
)

// GetValidFloatInput validates user input for numeric fields (e.g., price, amount).
func GetValidFloatInput(prompt string) float64 {
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
