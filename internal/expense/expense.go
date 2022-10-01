package expense

import "time"

// Expense represents a single expense at any given time
type Expense struct {
	ID          string
	Title       string
	Description string
	Price       float32
	At          time.Time
}

// NewExpense creates a new expense
func NewExpense(title, description string, price float32, at time.Time) *Expense {
	return &Expense{
		Title:       title,
		Description: description,
		Price:       price,
		At:          at,
	}
}
