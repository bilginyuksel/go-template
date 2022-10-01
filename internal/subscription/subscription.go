package subscription

import "time"

type Status string

const (
	Active   Status = "active"
	Canceled Status = "canceled"
)

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

type Settings struct {
	Notify     bool
	BeforeDays int
}

func (s *Subscription) NextNotice() time.Time {
	now := time.Now()
	payday := time.Date(now.Year(), now.Month(), s.MonthlyPayday, 0, 0, 0, 0, now.Location())
	return payday.Add(-time.Duration(s.Settings.BeforeDays) * 24 * time.Hour)
}
