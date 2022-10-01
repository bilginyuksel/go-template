package subscription

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type (
	// Repository is the interface that provides subscription storage methods
	Repository interface {
		Insert(ctx context.Context, subs *Subscription) (string, error)
		UpdateNoticeTime(ctx context.Context, id string, noticeAt time.Time) error
		Filter(ctx context.Context, filter Filter) ([]Subscription, error)
	}

	// LocalEventChannel is the interface that provides local event methods
	LocalEventChannel interface {
		Publish(ctx context.Context, event interface{}) error
	}
)

// Service provides subscription apis
type Service struct {
	repo Repository
	lec  LocalEventChannel
}

// NewService creates a new subscription service
func NewService(repo Repository, lec LocalEventChannel) *Service {
	return &Service{
		repo: repo,
		lec:  lec,
	}
}

// CreateSubscription creates a new subscription record
// Users configures set of settings for their subscription notifications
func (s *Service) CreateSubscription(ctx context.Context, subs *Subscription) (string, error) {
	subs.NoticeAt = subs.NextNotice()

	return s.repo.Insert(ctx, subs)
}

// CancelSubscription cancels a subscription
// Canceled subscriptions are not deleted, but are marked as canceled
// Notice and expense creation is stopped
func (s *Service) CancelSubscription(ctx context.Context, id string) error {
	return nil
}

// Filter is used to filter list of subscriptions
type Filter struct {
	Status   Status
	NoticeAt time.Time
}

// FilterSubscriptions filters subscriptions
func (s *Service) FilterSubscriptions(ctx context.Context, f Filter) ([]Subscription, error) {
	return s.repo.Filter(ctx, f)
}

// NotifySubscription is called when a subscription notice is received
// Sends a notification to the user
func (s *Service) NotifySubscription(ctx context.Context, subs *Subscription) error {
	if err := s.repo.UpdateNoticeTime(ctx, subs.ID, subs.NextNotice()); err != nil {
		zap.L().Error("update subscription next notice time failed", zap.Error(err))
		return err
	}

	// if err := s.publishExpense(ctx, subscription); err != nil {
	// 	return err
	// }

	// return s.repo.Update(ctx, subscription)
	return nil
}

// notificationEvent is the event that is published when a subscription is notified
type notificationEvent struct {
}

func (s *Service) publishNotificationEvent(ctx context.Context) {

}
