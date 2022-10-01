package subscription

import "time"

// Status is the subscription status
type Status string

const (
	// Active subscriptions will be notified when it's notice time
	Active Status = "active"
	// Canceled subscriptions will be kept to keep history, but will not be notified
	Canceled Status = "canceled"
)

// Subscription is the model for subscription
type Subscription struct {
	ID string

	Company     string
	Service     string
	Price       float32
	Description string
	Start       time.Time
	End         time.Time

	PaidInstallments int
	MonthlyPayday    int

	Settings Settings
	Status   Status

	NoticeAt time.Time
}

// Settings subscription notification settings
type Settings struct {
	Notify     bool
	BeforeDays int
}

// NextNotice returns the subscription next notice time
func (s *Subscription) NextNotice() time.Time {
	now := time.Now()
	payday := time.Date(now.Year(), now.Month(), s.MonthlyPayday, 0, 0, 0, 0, now.Location())
	return payday.Add(-time.Duration(s.Settings.BeforeDays) * 24 * time.Hour)
}
