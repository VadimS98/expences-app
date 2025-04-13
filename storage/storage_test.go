package storage

import (
	"expenses-app/models"
	"os"
	"testing"
)

func TestSaveAndLoadExpenses(t *testing.T) {
	filename := "test_expenses.json"
	defer os.Remove(filename) // Clean up after test

	expenses := []models.Expense{
		{Price: 10.5, Amount: 2, Category: "Food", Description: "Lunch"},
		{Price: 20, Amount: 1, Category: "Transport", Description: "Taxi"},
	}

	// Test saving
	err := SaveExpensesToFile(expenses, filename)
	if err != nil {
		t.Fatalf("Failed to save expenses: %v", err)
	}

	// Test loading
	loadedExpenses, err := LoadExpensesFromFile(filename)
	if err != nil {
		t.Fatalf("Failed to load expenses: %v", err)
	}

	if len(loadedExpenses) != len(expenses) {
		t.Errorf("Expected %d expenses, got %d", len(expenses), len(loadedExpenses))
	}
}
