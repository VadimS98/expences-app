package internal

import (
	"encoding/json"
	"expenses-app/models"
	"os"
)

// SaveExpenses saves expenses to a JSON file.
func SaveExpenses(expenses []models.Expense, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(expenses)
}

// LoadExpenses loads expenses from a JSON file.
func LoadExpenses(filename string) ([]models.Expense, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Expense{}, nil // Return empty slice if file doesn't exist.
		}
		return nil, err
	}
	defer file.Close()

	var expenses []models.Expense
	err = json.NewDecoder(file).Decode(&expenses)
	return expenses, err
}
