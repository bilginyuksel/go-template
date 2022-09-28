package subscription

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	Repository interface {
		Get(ctx context.Context, id string) (*Subscription, error)
		Insert(ctx context.Context, subs *Subscription) error
		Update(ctx context.Context, subs *Subscription) error
	}

	LocalEventChannel interface {
		Publish(ctx context.Context, event interface{}) error
	}
)

type Service struct {
	repo Repository
	lec  LocalEventChannel
}

// CreateSubscription creates a new subscription record
// Users configures set of settings for their subscription notifications
func (s *Service) CreateSubscription(ctx context.Context, subs Subscription) error {
	// create subscription record in db
	// create a timer to notify user when next payment is due
	// we need to get some settings about notifications by email and or how many days ahead to notify
	// how many paychecks to deduct from initial creation
	s.repo.Insert(ctx, &subs)
	return nil
}

// CancelSubscription cancels a subscription
// Canceled subscriptions are not deleted, but are marked as canceled
// Notice and expense creation is stopped
func (s *Service) CancelSubscription(ctx context.Context, id string) error {
	return nil
}

// ReceiveSubscriptionPaymentNotice is called when a subscription notice is received
// Checks subscription date and if it is due, creates an expense
// If it is not due, it updates the subscription date
// Sends a notification to the user
func (s *Service) ReceiveSubscriptionPaymentNotice(ctx context.Context, id string, days int) error {
	subscription, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}

	if subscription.IsDue() {
		if err := s.publishExpense(ctx, subscription); err != nil {
			return err
		}
	}

	subscription.UpdateSubscriptionState()

	return s.repo.Update(ctx, subscription)
}

type expenseEventMsg struct {
	Title       string
	Description string
	Price       float32
	Service     string
}

func (s *Service) publishExpense(ctx context.Context, subs *Subscription) error {
	msg := &expenseEventMsg{
		Title:       subs.Service,
		Description: fmt.Sprintf("Subscription for %s", subs.Service),
		Price:       subs.Price,
		Service:     subs.Service,
	}

	eventBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return s.lec.Publish(ctx, eventBytes)
}
