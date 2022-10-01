package expense

import (
	"context"
	"time"
)

// Repository is the interface that provides expense storage methods
type Repository interface {
	Insert(ctx context.Context, e *Expense) (string, error)
	Filter(ctx context.Context, f *Filter) ([]Expense, error)
}

// Service provides expense apis
type Service struct {
	repo Repository
}

// NewService creates a new expense service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateExpense creates a new expense
func (s *Service) CreateExpense(ctx context.Context, e *Expense) (string, error) {
	return s.repo.Insert(ctx, e)
}

// Filter is used to filter list of expenses
type Filter struct {
	TitleContains string

	LowerThanPrice  float32
	HigherThanPrice float32

	Before time.Time
	After  time.Time
}

// FilterExpenses filters expenses by given filter
func (s *Service) FilterExpenses(ctx context.Context, f *Filter) ([]Expense, error) {
	return s.repo.Filter(ctx, f)
}
