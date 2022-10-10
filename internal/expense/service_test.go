package expense_test

import (
	"context"
	"gotemplate/internal/expense"
	"testing"

	"gotemplate/internal/expense/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	mockRepository := mock.NewMockRepository(gomock.NewController(t))

	svc := expense.NewService(mockRepository)

	id, err := svc.CreateExpense(context.TODO(), &expense.Expense{})

	assert.NotNil(t, err)
	assert.Equal(t, "sfad", id)
}
