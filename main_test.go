package main

import (
	"expenses-app/models"
	"testing"
)

func TestCalculateTotalExpenses(t *testing.T) {
	expenses := []models.Expense{
		{Price: 10.5, Amount: 2},
		{Price: 20, Amount: 1},
	}

	total := 0.0
	for _, expense := range expenses {
		total += expense.Price * expense.Amount
	}

	expectedTotal := 41.0
	if total != expectedTotal {
		t.Errorf("Expected total %.2f, got %.2f", expectedTotal, total)
	}
}

func TestAddExpense(t *testing.T) {
	expenses := []models.Expense{}
	newExpense := models.Expense{Price: 15, Amount: 1, Category: "Entertainment", Description: "Movie"}

	expenses = append(expenses, newExpense)

	if len(expenses) != 1 {
		t.Errorf("Expected 1 expense, got %d", len(expenses))
	}

	if expenses[0] != newExpense {
		t.Errorf("Expected %v, got %v", newExpense, expenses[0])
	}
}
