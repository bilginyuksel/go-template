package subscription

import (
	"context"
	"encoding/json"
	"gotemplate/pkg/errors"
	"time"
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
		Publish(ctx context.Context, event string, msg []byte)
	}
)

// Service provides subscription apis
type Service struct {
	notificationEventChannel string

	repo Repository
	lec  LocalEventChannel
}

// NewService creates a new subscription service
func NewService(repo Repository, lec LocalEventChannel) *Service {
	return &Service{
		notificationEventChannel: "gotemplate.notification",

		repo: repo,
		lec:  lec,
	}
}

// CreateSubscription creates a new subscription record
// Users configures set of settings for their subscription notifications
func (s *Service) CreateSubscription(ctx context.Context, subs *Subscription) (string, error) {
	subs.NoticeAt = subs.NextNotice()

	id, err := s.repo.Insert(ctx, subs)
	return id, errors.Wrap(err, "subscription_service: insert subscription failed")
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
	subs, err := s.repo.Filter(ctx, f)
	return subs, errors.Wrap(err, "subscription_service: filter subscription failed")
}

// NotifySubscription is called when a subscription notice is received
// Sends a notification to the user
func (s *Service) NotifySubscription(ctx context.Context, subs *Subscription) error {
	if err := s.publishNotificationEvent(ctx, subs); err != nil {
		return errors.Wrap(err, "subscription_service: publish notification event failed")
	}
	return errors.Wrap(
		s.repo.UpdateNoticeTime(ctx, subs.ID, subs.NextNotice()),
		"subscription_service: update notice time failed",
	)
}

// notificationEvent is the event that is published when a subscription is notified
type notificationEvent struct {
	Company  string  `json:"company"`
	Service  string  `json:"service"`
	Price    float32 `json:"price"`
	DaysLeft int     `json:"daysLeft"`
}

func (s *Service) publishNotificationEvent(ctx context.Context, subs *Subscription) error {
	msg := notificationEvent{
		Company:  subs.Company,
		Service:  subs.Service,
		Price:    subs.Price,
		DaysLeft: subs.Settings.BeforeDays,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "subscription_service: marshal notification event msg failed")
	}

	s.lec.Publish(ctx, s.notificationEventChannel, msgBytes)

	return nil
}
