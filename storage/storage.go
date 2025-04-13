package storage

import (
	"encoding/json"
	"expenses-app/models"
	"os"
)

func SaveExpensesToFile(expenses []models.Expense, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(expenses)
}

func LoadExpensesFromFile(filename string) ([]models.Expense, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Expense{}, nil // Return empty list if file doesn't exist
		}
		return nil, err
	}
	defer file.Close()

	var expenses []models.Expense
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&expenses)
	return expenses, err
}
