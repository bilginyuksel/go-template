package main

import (
	"context"
	"fmt"
	"gotemplate/internal/subscription"
	subscription_adapter "gotemplate/internal/subscription/adapter"
	subscription_port "gotemplate/internal/subscription/port"
	"os"

	"github.com/labstack/echo/v4"
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

	subscriptionRepository := subscription_adapter.NewMongo(client.Database("gotemplate").Collection("subscriptions"))
	subscriptionService := subscription.NewService(subscriptionRepository, nil)
	subscriptionRestHandler := subscription_port.NewSubscriptionRestHandler(subscriptionService)

	e := echo.New()

	e.POST("/subscriptions", subscriptionRestHandler.CreateSubscription)
	e.GET("/subscriptions", subscriptionRestHandler.ListSubscriptions)

	if err := e.Start(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("shutting down the server", zap.Error(err))
	}
}
