package main

import (
	"context"
	"gotemplate/internal/expense"
	expense_adapter "gotemplate/internal/expense/adapter"
	"gotemplate/internal/subscription"
	subscription_adapter "gotemplate/internal/subscription/adapter"
	"gotemplate/pkg/broker"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main() {
	// Set global logger
	_ = zap.ReplaceGlobals(zap.NewExample())

	env := os.Getenv("APP_ENV")

	conf := readConfig(env)
	zap.L().Info("application started..", zap.Any("conf", conf))

	client := connectMongoClient(context.Background(), conf)
	b := broker.New()

	subscriptionService := newSubscriptionService(client, b)
	expenseService := newExpenseService(client)

	cronQuit := make(chan struct{}, 1)
	echoQuit := make(chan struct{}, 1)

	go runCronjob(conf, subscriptionService, cronQuit)
	go runEchoServer(conf, expenseService, subscriptionService, echoQuit)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	zap.L().Info("application shutting down..")

	// notify echo server to shutdown
	echoQuit <- struct{}{}

	// expect echo goroutine to notify us that it has stopped
	<-echoQuit

	disconnectMongoClient(context.Background(), client)
}

func connectMongoClient(ctx context.Context, conf Config) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.URI))
	if err != nil {
		panic(err)
	}

	return client
}

func disconnectMongoClient(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func newSubscriptionService(client *mongo.Client, b *broker.Broker) *subscription.Service {
	subscriptionRepository := subscription_adapter.NewMongo(client.Database("gotemplate").Collection("subscriptions"))

	return subscription.NewService(subscriptionRepository, b)
}

func newExpenseService(client *mongo.Client) *expense.Service {
	expenseRepository := expense_adapter.NewMongo(client.Database("gotemplate").Collection("expenses"))

	return expense.NewService(expenseRepository)
}
