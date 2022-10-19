package main

import (
	"context"
	"fmt"
	"gotemplate/internal/expense"
	expense_port "gotemplate/internal/expense/port"
	"gotemplate/internal/subscription"
	subscription_port "gotemplate/internal/subscription/port"
	"gotemplate/pkg/errors"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
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
	e.Use(middlewareEchoRequestID())

	e.GET("/health", healthcheck)

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	expenseRestHandler := expense_port.NewExpenseRestHandler(expenseService)
	subscriptionRestHandler := subscription_port.NewSubscriptionRestHandler(subscriptionService)

	e.HTTPErrorHandler = errors.EchoErrorHandler(
		expenseRestHandler.HandleErrors,
		subscriptionRestHandler.HandleErrors,
	)

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

func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

const ctxRequestID = "request_id"

func middlewareEchoRequestID() echo.MiddlewareFunc {
	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, id string) {
			//nolint:staticcheck
			contextWithRequestID := context.WithValue(c.Request().Context(), ctxRequestID, id)
			requestWithContext := c.Request().WithContext(contextWithRequestID)

			c.SetRequest(requestWithContext)
		},
	})
}
