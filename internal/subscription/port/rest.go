package port

import (
	"context"
	"gotemplate/internal/subscription"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, subs *subscription.Subscription) (string, error)
}

// SubscriptionRestHandler is responsible for handling subscription related requests
// It is a rest port for the application and uses echo library to provide restful api
type SubscriptionRestHandler struct {
	svc SubscriptionService
}

// NewSubscriptionRestHandler creates a new subscription rest handler
// SubscriptionRestHandler is responsible for handling subscription related requests
func NewSubscriptionRestHandler(svc SubscriptionService) *SubscriptionRestHandler {
	return &SubscriptionRestHandler{
		svc: svc,
	}
}

type (
	// createSubscriptionRequest is the request body for creating a new subscription
	createSubscriptionRequest struct {
		Company          string    `json:"company"`
		Service          string    `json:"service"`
		Price            float32   `json:"price"`
		Description      string    `json:"description"`
		Start            time.Time `json:"start"`
		End              time.Time `json:"end"`
		PaidInstallments int       `json:"paidInstallments"`
		MonthlyPayday    int       `json:"monthlyPayday"`
		Settings         struct {
			NoticeBeforeDays int  `json:"noticeBeforeDays"`
			Notify           bool `json:"notify"`
		} `json:"settings"`
	}

	// createSubscriptionResponse is the response body for creating a new subscription
	createSubscriptionResponse struct {
		ID string `json:"id"`
	}
)

func (c *createSubscriptionRequest) toSubscription() *subscription.Subscription {
	return &subscription.Subscription{
		Company:          c.Company,
		Service:          c.Service,
		Price:            c.Price,
		Description:      c.Description,
		Start:            c.Start,
		End:              c.End,
		PaidInstallments: c.PaidInstallments,
		MonthlyPayday:    c.MonthlyPayday,
		Settings: subscription.Settings{
			BeforeDays: c.Settings.NoticeBeforeDays,
			Notify:     c.Settings.Notify,
		},
	}
}

// CreateSubscription creates a new subscription record
func (h *SubscriptionRestHandler) CreateSubscription(c echo.Context) error {
	var req createSubscriptionRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	id, err := h.svc.CreateSubscription(c.Request().Context(), req.toSubscription())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createSubscriptionResponse{id})
}
