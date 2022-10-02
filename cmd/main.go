package main

import (
	"context"
	"fmt"
	"gotemplate/internal/expense"
	expense_adapter "gotemplate/internal/expense/adapter"
	expense_port "gotemplate/internal/expense/port"
	"gotemplate/internal/subscription"
	subscription_adapter "gotemplate/internal/subscription/adapter"
	subscription_port "gotemplate/internal/subscription/port"
	"gotemplate/pkg/broker"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	env := os.Getenv("APP_ENV")

	conf := readConfig(env)
	logger.Info("application started..", zap.Any("conf", conf))

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.Mongo.URI))
	if err != nil {
		panic(err)
	}

	b := broker.New()

	subscriptionRepository := subscription_adapter.NewMongo(client.Database("gotemplate").Collection("subscriptions"))
	subscriptionService := subscription.NewService(subscriptionRepository, b)
	subscriptionRestHandler := subscription_port.NewSubscriptionRestHandler(subscriptionService)
	subscriptionCronjob := subscription_port.NewSubscriptionNotificationCronjob(subscriptionService, 10)

	expenseRepository := expense_adapter.NewMongo(client.Database("gotemplate").Collection("expenses"))
	expenseService := expense.NewService(expenseRepository)
	expenseRestHandler := expense_port.NewExpenseRestHandler(expenseService)

	c := cron.New()
	if _, err := c.AddFunc("@every 10m", subscriptionCronjob.Notify); err != nil {
		panic(err)
	}
	go c.Run()
	defer c.Stop()

	e := echo.New()

	e.POST("/subscriptions", subscriptionRestHandler.CreateSubscription)
	e.GET("/subscriptions", subscriptionRestHandler.ListSubscriptions)

	e.POST("/expenses", expenseRestHandler.CreateExpense)
	e.GET("/expenses", expenseRestHandler.FilterExpenses)

	if err := e.Start(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("shutting down the server", zap.Error(err))
	}
}
