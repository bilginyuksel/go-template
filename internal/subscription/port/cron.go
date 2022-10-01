package port

import (
	"context"
	"gotemplate/internal/subscription"
	"time"

	"go.uber.org/zap"
)

type SubscriptionNotifier interface {
	FilterSubscriptions(ctx context.Context, f subscription.Filter) ([]subscription.Subscription, error)
	NotifySubscription(ctx context.Context, subs *subscription.Subscription) error
}

type SubscriptionNotificationCronjob struct {
	maxNumOfConcurrentGoroutines int

	svc SubscriptionNotifier
}

func NewSubscriptionNotificationCronjob(svc SubscriptionNotifier, maxNumOfConcurrentGoroutines int) *SubscriptionNotificationCronjob {
	return &SubscriptionNotificationCronjob{
		svc:                          svc,
		maxNumOfConcurrentGoroutines: maxNumOfConcurrentGoroutines,
	}
}

// Notify runs the cronjob to notify subscriptions
func (c *SubscriptionNotificationCronjob) Notify() {
	ctx := context.Background()

	dueSubscriptions, err := c.svc.FilterSubscriptions(ctx, subscription.Filter{
		Status:   subscription.Active,
		NoticeAt: time.Now(),
	})
	if err != nil {
		zap.L().Error("failed to filter subscriptions", zap.Error(err))
		return
	}

	c.notifyAllDueSubscriptions(ctx, dueSubscriptions)
}

func (c *SubscriptionNotificationCronjob) notifyAllDueSubscriptions(ctx context.Context, dueSubscriptions []subscription.Subscription) {
	routines := make(chan struct{}, c.maxNumOfConcurrentGoroutines)

	for _, subs := range dueSubscriptions {
		routines <- struct{}{}

		go func(subs *subscription.Subscription) {
			if err := c.svc.NotifySubscription(ctx, subs); err != nil {
				zap.L().Error("failed to notify subscription", zap.Error(err))
			}

			<-routines
		}(&subs)
	}
}
