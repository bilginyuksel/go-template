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

	NoticeAt      time.Time
	NextPaymentAt time.Time
	LastPaidAt    time.Time
}

type Settings struct {
	Notify     bool
	BeforeDays int
}

func (s *Subscription) CalculateSubscriptionNotice() {
	daysBefore := s.Settings.BeforeDays

	now := time.Now()
	nextPayday := time.Date(now.Year(), now.Month(), s.PaymentDayOfMonth, 0, 0, 0, 0, now.Location())
	s.NoticeAt = nextPayday.Add(-time.Hour * 24 * time.Duration(daysBefore))
}

func (s *Subscription) UpdateSubscriptionState() {
	s.LastPaidAt = s.NextPaymentAt
	s.NextPaymentAt = s.NextPaymentAt.Add(_durationMonth)
}

func (s *Subscription) IsDue() bool {
	return s.NextPaymentAt.Before(time.Now())
}
