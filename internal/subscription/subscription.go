package subscription

import "time"

const _durationMonth = time.Hour * 24 * 30

type Period string

const (
	PeriodWeekly  Period = "weekly"
	PeriodMonthly Period = "monthly"
	PeriodYearly  Period = "yearly"
)

type Subscription struct {
	ID string

	PaymentAccount string

	Company     string
	Service     string
	Period      Period
	Price       float32
	Description string
	Start       time.Time
	End         time.Time

	PaymentDayOfMonth int

	// Settings for notifications
	Settings Settings

	NextPaymentAt time.Time
	LastPaidAt    time.Time
}

type Settings struct {
	Notify     bool
	BeforeDays int
}

func (s *Subscription) UpdateSubscriptionState() {
	s.LastPaidAt = s.NextPaymentAt
	s.NextPaymentAt = s.NextPaymentAt.Add(_durationMonth)
}

func (s *Subscription) IsDue() bool {
	return s.NextPaymentAt.Before(time.Now())
}
