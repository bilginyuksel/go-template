package port

import (
	"context"
	"gotemplate/internal/expense"
	"gotemplate/pkg/errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// ExpenseService is the business implementation of expense
type ExpenseService interface {
	CreateExpense(ctx context.Context, e *expense.Expense) (string, error)
	FilterExpenses(ctx context.Context, f *expense.Filter) ([]expense.Expense, error)
}

// ExpenseRestHandler is the interface that provides expense rest apis
type ExpenseRestHandler struct {
	svc ExpenseService
}

// NewExpenseRestHandler creates a new expense rest handler
func NewExpenseRestHandler(svc ExpenseService) *ExpenseRestHandler {
	return &ExpenseRestHandler{
		svc: svc,
	}
}

// HandleErrors handle expense rest errors
func (h *ExpenseRestHandler) HandleErrors(err error) *errors.HTTPError {
	return nil
}

type (
	createExpenseRequest struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Price       float32   `json:"price"`
		At          time.Time `json:"at"`
	}

	createExpenseResponse struct {
		ID string `json:"id"`
	}
)

func (c *createExpenseRequest) toExpense() *expense.Expense {
	return &expense.Expense{
		Title:       c.Title,
		Description: c.Description,
		Price:       c.Price,
		At:          c.At,
	}
}

// CreateExpense creates a new expense
func (h *ExpenseRestHandler) CreateExpense(c echo.Context) error {
	var req createExpenseRequest
	if err := c.Bind(&req); err != nil {
		zap.L().Error("failed to bind request", zap.Error(err))
		return err
	}

	id, err := h.svc.CreateExpense(c.Request().Context(), req.toExpense())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createExpenseResponse{id})
}

type expenseResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	At          time.Time `json:"at"`
}

func newExpenseResponse(e *expense.Expense) *expenseResponse {
	return &expenseResponse{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		Price:       e.Price,
		At:          e.At,
	}
}

type (
	filterExpenseRequest struct {
		TitleContains string    `json:"titleContains"`
		LowerThan     float32   `json:"lowerThan"`
		HigherThan    float32   `json:"higherThan"`
		Before        time.Time `json:"before"`
		After         time.Time `json:"after"`
	}

	filterExpenseResponse []expenseResponse
)

func (req *filterExpenseRequest) toFilter() *expense.Filter {
	return &expense.Filter{
		TitleContains:   req.TitleContains,
		LowerThanPrice:  req.LowerThan,
		HigherThanPrice: req.HigherThan,
		Before:          req.Before,
		After:           req.After,
	}
}

// FilterExpenses filters expenses by given filter
func (h *ExpenseRestHandler) FilterExpenses(c echo.Context) error {
	var req filterExpenseRequest
	if err := c.Bind(&req); err != nil {
		zap.L().Error("failed to bind request", zap.Error(err))
		return err
	}

	expenses, err := h.svc.FilterExpenses(c.Request().Context(), req.toFilter())
	if err != nil {
		return err
	}

	var res filterExpenseResponse
	for idx := range expenses {
		res = append(res, *newExpenseResponse(&expenses[idx]))
	}

	return c.JSON(http.StatusOK, res)
}
