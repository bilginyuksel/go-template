package main

import (
	"fmt"
	"gotemplate/internal/expense"
	expense_port "gotemplate/internal/expense/port"
	"gotemplate/internal/subscription"
	subscription_port "gotemplate/internal/subscription/port"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func runEchoServer(
	conf Config,
	expenseService *expense.Service,
	subscriptionService *subscription.Service,
) {

	e := echo.New()

	expenseRestHandler := expense_port.NewExpenseRestHandler(expenseService)
	subscriptionRestHandler := subscription_port.NewSubscriptionRestHandler(subscriptionService)

	e.POST("/expenses", expenseRestHandler.CreateExpense)
	e.GET("/expenses", expenseRestHandler.FilterExpenses)

	e.POST("/subscriptions", subscriptionRestHandler.CreateSubscription)
	e.GET("/subscriptions", subscriptionRestHandler.ListSubscriptions)

	if err := e.Start(fmt.Sprintf(":%d", conf.Port)); err != nil {
		zap.L().Fatal("shutting down the server", zap.Error(err))
	}
}
