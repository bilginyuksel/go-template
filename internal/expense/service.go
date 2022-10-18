package expense

import (
	"context"
	"gotemplate/pkg/errors"
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
	id, err := s.repo.Insert(ctx, e)
	return id, errors.Wrap(err, "expense_service: insert expense failed")
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
	expenses, err := s.repo.Filter(ctx, f)
	return expenses, errors.Wrap(err, "expense_service: filter expenses failed")
}
