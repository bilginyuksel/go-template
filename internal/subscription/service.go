package subscription

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type (
	Repository interface {
		Insert(ctx context.Context, subs *Subscription) (string, error)
		Filter(ctx context.Context, filter Filter) ([]Subscription, error)
	}

	LocalEventChannel interface {
		Publish(ctx context.Context, event interface{}) error
	}
)

type Service struct {
	repo Repository
	lec  LocalEventChannel
}

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

// ListSubscriptions lists all subscriptions
func (s *Service) FilterSubscriptions(ctx context.Context, f Filter) ([]Subscription, error) {
	return s.repo.Filter(ctx, f)
}

// NotifySubscription is called when a subscription notice is received
// Checks subscription date and if it is due, creates an expense
// If it is not due, it updates the subscription date
// Sends a notification to the user
func (s *Service) NotifySubscription(ctx context.Context, subs *Subscription) error {
	zap.L().Info("notifying subscription", zap.String("id", subs.ID))
	// if err := s.publishExpense(ctx, subscription); err != nil {
	// 	return err
	// }

	// return s.repo.Update(ctx, subscription)
	return nil
}

// type expenseEventMsg struct {
// 	Title       string
// 	Description string
// 	Price       float32
// 	Service     string
// }

// func (s *Service) publishExpense(ctx context.Context, subs *Subscription) error {
// 	msg := &expenseEventMsg{
// 		Title:       subs.Service,
// 		Description: fmt.Sprintf("Subscription for %s", subs.Service),
// 		Price:       subs.Price,
// 		Service:     subs.Service,
// 	}

// 	eventBytes, err := json.Marshal(msg)
// 	if err != nil {
// 		return err
// 	}

// 	return s.lec.Publish(ctx, eventBytes)
// }
