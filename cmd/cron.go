package main

import (
	"gotemplate/internal/subscription"
	subscription_port "gotemplate/internal/subscription/port"

	"github.com/robfig/cron/v3"
)

func runCronjob(
	conf Config,
	subscriptionService *subscription.Service,
) {
	c := cron.New()

	subscriptionCronjob := subscription_port.NewSubscriptionNotificationCronjob(subscriptionService, 10)

	if _, err := c.AddFunc("@every 10m", subscriptionCronjob.Notify); err != nil {
		panic(err)
	}

	c.Run()
	defer c.Stop()
}
