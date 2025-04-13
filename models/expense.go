package models

type Expense struct {
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}
