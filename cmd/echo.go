package main

import (
	"context"
	"fmt"
	"gotemplate/internal/expense"
	expense_port "gotemplate/internal/expense/port"
	"gotemplate/internal/subscription"
	subscription_port "gotemplate/internal/subscription/port"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func runEchoServer(
	conf Config,
	expenseService *expense.Service,
	subscriptionService *subscription.Service,
	quit chan struct{},
) {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	expenseRestHandler := expense_port.NewExpenseRestHandler(expenseService)
	subscriptionRestHandler := subscription_port.NewSubscriptionRestHandler(subscriptionService)

	e.POST("/expenses", expenseRestHandler.CreateExpense)
	e.GET("/expenses", expenseRestHandler.FilterExpenses)

	e.POST("/subscriptions", subscriptionRestHandler.CreateSubscription)
	e.GET("/subscriptions", subscriptionRestHandler.ListSubscriptions)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", conf.Port)); err != nil {
			zap.L().Fatal("shutting down the server", zap.Error(err))
		}
	}()

	<-quit

	zap.L().Info("shutting down the server..")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		zap.L().Error("could not gracefully shutdown the server", zap.Error(err))
	}

	// notify caller that http server has stopped
	quit <- struct{}{}
}
