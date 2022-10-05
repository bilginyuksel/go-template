package main

import (
	"gotemplate/internal/subscription"
	subscription_port "gotemplate/internal/subscription/port"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func runCronjob(
	conf Config,
	subscriptionService *subscription.Service,
	quit chan struct{},
) {
	c := cron.New()

	subscriptionCronjob := subscription_port.NewSubscriptionNotificationCronjob(subscriptionService, 10)

	if _, err := c.AddFunc("@every 10m", subscriptionCronjob.Notify); err != nil {
		panic(err)
	}

	go c.Run()

	// wait for quit signal
	<-quit

	ctx := c.Stop()

	// wait for cronjob to stop§
	<-ctx.Done()

	zap.L().Info("cronjob stopped")

	// notify caller that cron job has stopped
	quit <- struct{}{}
}
