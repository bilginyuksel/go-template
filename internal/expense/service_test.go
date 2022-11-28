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
	expectedExpense := &expense.Expense{}

	mockRepository := mock.NewMockRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		Insert(gomock.Any(), expectedExpense).
		Return("expected-id", nil)

	svc := expense.NewService(mockRepository)

	id, err := svc.CreateExpense(context.TODO(), expectedExpense)

	assert.NoError(t, err)
	assert.Equal(t, "expected-id", id)
}
