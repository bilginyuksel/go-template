package main

import (
	"context"
	"gotemplate/internal/expense"
	expense_adapter "gotemplate/internal/expense/adapter"
	"gotemplate/internal/subscription"
	subscription_adapter "gotemplate/internal/subscription/adapter"
	"gotemplate/pkg/broker"
	"os"

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

	client := connectMongoClient(context.Background(), conf)

	b := broker.New()

	subscriptionService := newSubscriptionService(client, b)
	expenseService := newExpenseService(client)

	runCronjob(conf, subscriptionService)
	runEchoServer(conf, expenseService, subscriptionService)
}

func connectMongoClient(ctx context.Context, conf Config) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.URI))
	if err != nil {
		panic(err)
	}

	return client
}

func newSubscriptionService(client *mongo.Client, b *broker.Broker) *subscription.Service {
	subscriptionRepository := subscription_adapter.NewMongo(client.Database("gotemplate").Collection("subscriptions"))

	return subscription.NewService(subscriptionRepository, b)
}

func newExpenseService(client *mongo.Client) *expense.Service {
	expenseRepository := expense_adapter.NewMongo(client.Database("gotemplate").Collection("expenses"))

	return expense.NewService(expenseRepository)
}
